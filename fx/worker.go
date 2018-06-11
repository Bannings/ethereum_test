package fx

import (
	"time"

	"github.com/eddyzhou/log"
)

func ExcuteTransaction(done <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Infof("Process for data to blockchain is done! time:%v", time.Now())
			return
		case <-ticker.C:
			unfinsherdCmds := getUnfinshedCmd()
			supplierMap := sectionalizeCmd(unfinsherdCmds)
			for _, cmds := range supplierMap {
				go run(done, cmds)
			}
		}
	}
}

func run(done <-chan struct{}, cmds []Command) {
	for r := range storeToEth(done, gen(done, cmds...)) {
		if r.err != nil {
			log.Errorf("execute transaction failed: id: %v, err: %v", r.Id, r.err)
		}
	}
}

func gen(done <-chan struct{}, cmds ...Command) <-chan Command {
	out := make(chan Command)
	go func() {
		for _, cmd := range cmds {
			select {
			case out <- cmd:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

func storeToEth(done <-chan struct{}, in <-chan Command) <-chan ProcessResult {
	out := make(chan ProcessResult)
	go func() {
		defer close(out)
		for cmd := range in {
			err := DefaultExecutor().Execute(cmd)
			select {
			case out <- ProcessResult{Id: cmd.Tx.Id, err: err}:
			case <-done:
				return
			}
		}
	}()
	return out
}

func getUnfinshedCmd() []Command {
	e := DefaultExecutor()
	p := &CmdProcessor{db: e.db}
	cmds, err := p.getCmd()
	if err != nil {
		log.Errorf("Get unfinshed cmds fail:%v", err)
		return nil
	}
	return cmds
}

func sectionalizeCmd(cmds []Command) map[string][]Command {
	supplierMap := make(map[string][]Command)

	for _, cmd := range cmds {

		if cmd.Tx.TxType == MintFX {
			if _, ok := supplierMap["cf"]; ok {
				temp := supplierMap["cf"]
				temp = append(temp, cmd)
				supplierMap["cf"] = temp
			} else {
				supplierMap["cf"] = []Command{}
				supplierMap["cf"] = append(supplierMap["cf"], cmd)
			}
		} else {
			if _, ok := supplierMap[cmd.Tx.Input[0].Owner]; ok {
				supplierMap[cmd.Tx.Input[0].Owner] = append(supplierMap[cmd.Tx.Input[0].Owner], cmd)
			} else {
				temp := []Command{cmd}
				supplierMap[cmd.Tx.Input[0].Owner] = temp
			}
		}
	}

	return supplierMap
}
