package iron

import (
	"bufio"
    "os"
    "github.com/jkbbwr/iron/iron/sys"
    . "github.com/jkbbwr/iron/iron/parser"
    . "github.com/jkbbwr/iron/iron/types"
)

const NUM_REGISTERS int = 12000

type VM struct {
	Halted      bool
	Register    []FeType
	Locals      []FeType
	Window      []FeType
	Program     Program
	System      sys.System
	JumpTable   map[string]int
	Initialized bool
	Pc          int
	Wp          int
	Cf          int
}

func NewVM() *VM {
	registers := make([]FeType, NUM_REGISTERS)
	locals := make([]FeType, 10)
	for i := range locals {
		locals[i] = FeNothing{}
	}
	for i := range registers {
		registers[i] = FeNothing{}
	}

	vm := VM{
		Initialized: false,
		System:      sys.NewSystemMapper(),
		Locals:      locals,
		Register:    registers,
		Window:      registers[0:25],
        JumpTable:   make(map[string]int),
		Pc:          0,
		Wp:          0,
	}
	return &vm
}

func (vm *VM) Load(path string) {
	parser := NewParser(path)
	vm.Program = parser.Parse()
	vm.Initialized = true
	log.Info("Program parsed and loaded.")
}

func (vm *VM) IncrWindow() {
	vm.Wp += 17
	vm.Window = vm.Register[vm.Wp : vm.Wp+25]
}

func (vm *VM) DecrWindow() {
	vm.Wp -= 17
	vm.Window = vm.Register[vm.Wp : vm.Wp+25]
}

func (vm *VM) Step() {
	if vm.Pc >= vm.Program.CodeLength {
		log.Critical("PC is greater than the length of instructions. Halting now.")
		vm.Exit()
	}

	instr := vm.Program.Instructions[vm.Pc]
	vm.Pc++
	//log.Debug("Opcode = %s", debugOpcode(instr.OpCode))
	switch instr.OpCode {
	case LConst:
		vm.LConst(instr)
	case LLocal:
		vm.LLocal(instr)
	case SLocal:
		vm.SLocal(instr)
	case Move:
		vm.Move(instr)
	case Add:
		vm.Add(instr)
	case Sub:
		vm.Sub(instr)
	case Div:
		vm.Div(instr)
	case Mul:
		vm.Mul(instr)
	case Mod:
		vm.Mod(instr)
	case Cmp:
		vm.Cmp(instr)
	case Inc:
		vm.Inc(instr)
	case Dec:
		vm.Dec(instr)
	case NewList:
		vm.NewList(instr)
	case AddList:
		vm.AddList(instr)
	case PopList:
		vm.PopList(instr)
	case JmpEq:
		vm.JmpCmp(instr, Equal)
	case JmpNEq:
		vm.JmpCmp(instr, NotEqual)
	case JmpGt:
		vm.JmpCmp(instr, GreaterThan)
	case JmpGtEq:
		vm.JmpCmp(instr, GreaterThanOrEqual)
	case JmpLt:
		vm.JmpCmp(instr, LessThan)
	case JmpLtEq:
		vm.JmpCmp(instr, LessThanOrEqual)
	case Call:
		vm.Call(instr)
	case Sys:
		vm.Sys(instr)
	case Ret:
		vm.Ret(instr)
	case Halt:
		vm.Exit()
	}
}

func (vm *VM) Run() {
	if !vm.Initialized {
		panic("No program loaded. Call .Load(sourcePath)!")
	}
	for {
		vm.Step()
		if vm.Halted {
			return
		}
	}

}

func (vm *VM) RunStep() {
	if !vm.Initialized {
		panic("No program loaded. Call .Load(sourcePath)!")
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		reader.ReadLine()
		vm.Step()
		dumpWindow(vm)
		if vm.Halted {
			return
		}
	}
}

func (vm *VM) RunDebug() {
	if !vm.Initialized {
		panic("No program loaded. Call .Load(sourcePath)!")
	}
	for {
		vm.Debugger()
		vm.Step()

		if vm.Halted {
			return
		}
	}
}

func (vm *VM) Exit() {
	log.Info("Shutting down FeVM")
	vm.Halted = true
}
