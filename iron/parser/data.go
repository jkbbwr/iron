package parser

import . "github.com/jkbbwr/iron/iron/types"

const (
	LConst = iota
	LLocal
	SLocal
	Move
	Clr
	Nop
	Jmp
	Cmp
	Add
	Sub
	Div
	Mul
	Inc
	Dec
	Mod
	NewList
	AddList
	PopList
	Call
	Ret
	JmpEq
	JmpNEq
	JmpGt
	JmpGtEq
	JmpLt
	JmpLtEq
	Sys
	Halt = 255
)

const (
	Int = iota
	String
	Bool
	Nothing
)

type Instr struct {
	OpCode int
	Args   []int
}

type Func struct {
	Name         string
	Instructions InstrList
}

type Const struct {
	Type  int
	Value interface{}
}

type ConstList []FeType
type InstrList []Instr
type Functions map[string]Func

type Program struct {
	Name         string // should just always be "main"
	Consts       ConstList
	Funcs        Functions
	Instructions InstrList
	CodeLength   int // Stop it being computed constantly
}
