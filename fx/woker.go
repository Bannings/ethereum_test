package fx

import (
	"time"

	"github.com/eddyzhou/log"
)

// run
// ```
// done := make(chan struct{})
// defer close(done)
// go run(done)
// ```
// in main.go
func Run(done <-chan struct{}) {
	ticker := time.NewTicker(300 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Println("Done!")
			return
		case <-ticker.C:
			var cmds []Command
			// TODO: fetch cmd from db(transactions left join tx_procedure)
			for r := range storeToEth(done, gen(done, cmds...)) {
				if r.err != nil {
					log.Errorf("execute command failed: id: %v, err: %v", r.Id, r.err)
				}
			}
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
			_, err := DefaultExecutor().Execute(cmd)
			select {
			case out <- ProcessResult{Id: cmd.Tx.Id, err: err}:
			case <-done:
				return
			}
		}
	}()
	return out
}
