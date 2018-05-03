package fx

import (
	"math/big"
	"testing"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

func TestClosure(t *testing.T) {
	var consumedNum int
	tokenIds := make([]*big.Int, 3)
	packing := func() {
		jobId := new(big.Int).SetUint64(uint64(consumedNum / 3))
		t.Logf("jobId: %v", jobId)
		t.Logf("tokens: %+v", tokenIds)
	}

	var actualAmount uint64
	tokens := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, t := range tokens {
		actualAmount += 1
		consumedNum += 1
		i := idx % 3
		tokenIds[i] = new(big.Int).SetInt64(t)
		if actualAmount%3 == 0 {
			packing()

			size := len(tokens) - consumedNum
			if size > 3 {
				size = 3
			}
			tokenIds = make([]*big.Int, size)
		}
	}

	if actualAmount%3 > 0 {
		packing()
	}
}

func TestDB(t *testing.T) {
	tx := Transaction{Id: 1}
	cmd := Command{Tx: tx, startNonce: 1, receipts: make(map[string]*ethTypes.Receipt)}
	p := &CmdProcessor{cmd: cmd, db: db}
	if err := p.createProcedure(); err != nil {
		t.Fatal(err)
	}

	r := &ethTypes.Receipt{Status: 0}
	if err := p.updateReceipt(r); err != nil {
		t.Fatal(err)
	}

	r.Status = 1
	p.cmd.currNonce += 1
	if err := p.updateReceipt(r); err != nil {
		t.Fatal(err)
	}

	if err := p.finishProcedure(); err != nil {
		t.Fatal(err)
	}
}
