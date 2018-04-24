package fx

import (
	"fmt"
	"math/big"
	"strings"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type TxType uint8

const (
	Payment TxType = iota
	Discount
	SplitFX
	MintFX
	Confirm
)

var types = []string{
	"Payment",
	"Discount",
	"SplitFX",
	"MintFX",
	"Confirm",
}

func ParseType(typ string) (TxType, error) {
	switch strings.ToLower(typ) {
	case "payment":
		return Payment, nil
	case "discount":
		return Discount, nil
	case "splitfx":
		return SplitFX, nil
	case "mintfx":
		return MintFX, nil
	case "confirm":
		return Confirm, nil
	}

	var t TxType
	return t, fmt.Errorf("err type: %q", typ)
}

// -------------

type State uint8

const (
	Frozen State = iota
	Normal
	Historic
)

var states = []string{
	"Frozen",
	"Normal",
	"Historic",
}

func ParseState(state string) (State, error) {
	switch strings.ToLower(state) {
	case "frozen":
		return Frozen, nil
	case "normal":
		return Normal, nil
	case "historic":
		return Historic, nil
	}

	var s State
	return s, fmt.Errorf("err State: %q", state)
}

// ----------------

type Token struct {
	Id         big.Int
	Amount     uint64
	Owner      string // companyID
	State      State
	ExpireTime int64
}

type Transaction struct {
	Id     uint
	Input  []Token
	Output []Token
	TxId   uint64
	TxType TxType
}

func (t *Transaction) Sponsor() string {
	return t.Input[0].Owner
}

type Command struct {
	Tx         Transaction
	startNonce uint64
	currNonce  uint64
	receipts   map[string]*ethTypes.Receipt // key: string(nonce)
}
