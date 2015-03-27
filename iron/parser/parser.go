package parser

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"

    "github.com/op/go-logging"
    . "github.com/jkbbwr/iron/iron/types"
)

var log = logging.MustGetLogger("FeVM")

type Parser struct {
	Path   string
	Reader *bufio.Reader
}

func NewParser(path string) Parser {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return Parser{
		Path:   path,
		Reader: bufio.NewReader(fi),
	}
}

func (parser Parser) ReadString() string {
	var size uint32
	err := binary.Read(parser.Reader, binary.BigEndian, &size)
	if err != nil {
		panic(err)
	}
	bytes := make([]byte, int(size))
	length, err := io.ReadFull(parser.Reader, bytes)
	if err != nil {
		panic(err)
	}

	return string(bytes[:length])
}

func (parser Parser) ReadInt() int {
	var i int32
	err := binary.Read(parser.Reader, binary.BigEndian, &i)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func (parser Parser) ReadByte() int {
	var b uint8
	if err := binary.Read(parser.Reader, binary.BigEndian, &b); err != nil {
		panic(err)
	} else {
		return int(b)
	}
}

func (parser Parser) ParseConstList() ConstList {
	size := parser.ReadByte()

	consts := make(ConstList, size)
	for i := 0; i < size; i++ {
		consts[i] = parser.ParseConst()
	}
	return consts
}

func (parser Parser) ParseConst() FeType {
	switch parser.ReadByte() {
	case Bool:
		return FeBool(parser.ReadByte() != 0)
	case String:
		return FeString(parser.ReadString())
	case Nothing:
		return FeNothing{}
	case Int:
		return FeInt(parser.ReadInt())
	}
	return FeNothing{}
}

func (parser Parser) ParseInstrList() InstrList {
	size := parser.ReadByte()

	instrs := make(InstrList, size)
	for i := 0; i < size; i++ {
		instrs[i] = parser.ParseInstr()
	}
	return instrs
}

func (parser Parser) ParseInstr() Instr {
	instrType := parser.ReadByte()

	var args []int
	switch instrType {
	case Halt, Nop, Ret:
		break
	case LConst, LLocal, SLocal, Move, Add, Sub, Mul,
		Div, Mod, Cmp, AddList, PopList:
		args = []int{parser.ReadByte(), parser.ReadByte()}
	case Clr, Jmp, JmpEq, JmpNEq, JmpGt, JmpGtEq, JmpLt, JmpLtEq, Inc, Dec, Call, NewList, Sys:
		args = []int{parser.ReadByte()}
	}
	return Instr{instrType, args}
}

func (parser Parser) ParseFunc() Func {
	name := parser.ReadString()
	instructions := parser.ParseInstrList()
	return Func{name, instructions}
}

func (parser Parser) ParseFunctions() Functions {
	size := parser.ReadByte()
	funcs := make(Functions, size)
	for i := 0; i < size; i++ {
		function := parser.ParseFunc()
		funcs[function.Name] = function
	}
	return funcs
}

func (parser Parser) Parse() Program {
	main := parser.ReadString()

	consts := parser.ParseConstList()
	functions := parser.ParseFunctions()
	instructions := parser.ParseInstrList()
	return Program{main, consts, functions, instructions, len(instructions)}
}
