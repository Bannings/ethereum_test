package tokens

import (
	"time"

	"database/sql"
	"encoding/json"
	"github.com/eddyzhou/log"
	"gitlab.zhuronglink.com/chaincore/r2/g"
	"sync"
)

var (
	defaultDBConnection *sql.DB
	once                sync.Once
)

func ExcuteTransaction(done <-chan struct{}) {
	ticker := time.NewTicker(30 * time.Second)
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
				run(done, cmds)
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
	cmds, err := getCmd()
	if err != nil {
		log.Errorf("Get unfinshed cmds fail:%v", err)
		return nil
	}
	return cmds
}

func sectionalizeCmd(cmds []Command) map[string][]Command {
	supplierMap := make(map[string][]Command)

	for _, cmd := range cmds {
		txtype, _ := ParseType(cmd.Tx.TxType)
		if txtype == MintToken {
			if _, ok := supplierMap["admin"]; ok {
				temp := supplierMap["admin"]
				temp = append(temp, cmd)
				supplierMap["admin"] = temp
			} else {
				supplierMap["admin"] = []Command{}
				supplierMap["admin"] = append(supplierMap["admin"], cmd)
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

func getCmd() ([]Command, error) {
	var cmds []Command
	rows, err := DefaultDBConnection().Query("select * from( select t1.id,input,output,t1.state,t2.state process_state from transactions t1 left join cmd_procedure t2 on t1.id=transaction_id order by t1.created_at) as t3 where process_state is null;")
	//rows, err := p.db.Query("select deal_id,input,output,state from transactions where status='unprocessed' or status='failed' order by created_at")
	if err != nil {
		log.Error("fail to query unfinished transactions: %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var inputStr, outputStr, state, process_state string
		var input, output []Token
		var id uint

		rows.Scan(&id, &inputStr, &outputStr, &state, &process_state)
		json.Unmarshal([]byte(inputStr), &input)
		json.Unmarshal([]byte(outputStr), &output)
		tx := &Transaction{Id: id, Input: input, Output: output, TxType: state}
		cmd := &Command{Tx: *tx, txHashes: make(map[string]string)}
		cmds = append(cmds, *cmd)
	}
	return cmds, nil
}

func DefaultDBConnection() *sql.DB {
	once.Do(func() {
		conf := g.GetConfig()
		db, err := g.OpenDB(conf.DbConfig)
		if err != nil {
			panic(err)
		}
		defaultDBConnection = db
	})

	return defaultDBConnection
}
