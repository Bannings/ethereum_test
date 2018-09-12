package tokens

import (
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

type TxType uint8

const (
	Payment TxType = iota
	Setting
	SplitToken
	MintToken
	Cancellation
)

var types = []string{
	"Payment",
	"Setting",
	"SplitToken",
	"MintToken",
	"Cancellation",
}

func ParseType(typ string) (TxType, error) {
	switch strings.ToLower(typ) {
	case "payment":
		return Payment, nil
	case "setting":
		return Setting, nil
	case "splittoken":
		return SplitToken, nil
	case "minttoken":
		return MintToken, nil
	case "cancellation":
		return Cancellation, nil
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
	Id         big.Int `json:"id"`
	ParentId   big.Int `json:parentId`
	Amount     uint64  `json:"amount"`
	Owner      string  `json:"owner"` //company ID
	State      string  `json:"state"`
	ExpireTime int64   `json:"expire_time"`
}

type Transaction struct {
	Id     uint    `json:"id"`
	Input  []Token `json:"input"`
	Output []Token `json:"output"`
	TxId   string  `json:"tx_id"`
	TxType string  `json:"tx_type"`
}

type Management struct {
	Tokens []Token `json:"tokens"`
	Type   string  `json:"type"`
}

func (t *Transaction) Sponsor() string {
	return t.Input[0].Owner
}

type Command struct {
	Tx         Transaction
	startNonce uint64
	currNonce  uint64
	txHashes   map[string]string // key: string(nonce)
}

type ProcessResult struct {
	Id       uint
	Tx       *ethTypes.Transaction
	Supplier string
	err      error
}
