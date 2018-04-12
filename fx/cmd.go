package fx

import (
	"fmt"
	"math/big"
	"strings"
)

type CmdType uint8

const (
	Payment CmdType = iota
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

func ParseType(typ string) (CmdType, error) {
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

	var t CmdType
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
	ID         big.Int
	Amount     uint64
	Owner      string // companyID
	State      State
	ExpireTime int64
}

type Transaction struct {
	Input  []Token
	Output []Token
	TxId   uint64
}

func (t *Transaction) Sponsor() string {
	return t.Input[0].Owner
}

type Command struct {
	T    Transaction
	Type CmdType
}

// -------------

type ContractType uint8

const (
	CreateFX ContractType = iota
	CreateBox
	Transfer
	Split
)

// Instruction corresponds a call to an Ethernet contract
type Instruction struct {
	From  string
	To    string
	Nonce uint64
	Input []string
	Type  ContractType
}
