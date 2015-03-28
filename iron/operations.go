package iron

import (
	. "github.com/jkbbwr/iron/iron/parser"
	. "github.com/jkbbwr/iron/iron/types"
)

func (vm *VM) LConst(instruction Instr) {
	/*
	   LOADC dest const_num
	   Loading the constant from the pool into dest in the current window
	*/
	a := instruction.Args[0]
	b := instruction.Args[1]
	c := vm.Program.Consts[b]
	log.Debug("LConst(%+#v) -> R(%d)", c, a)
	vm.Window[a] = c
}

func (vm *VM) LLocal(instruction Instr) {
	reg := instruction.Args[0]
	local := instruction.Args[1]
	vm.Window[reg] = vm.Locals[local]
}

func (vm *VM) SLocal(instruction Instr) {
	index := instruction.Args[0]
	reg := instruction.Args[1]
	// We don't want locals to be too big but we might have lots and lots of locals.
	// If the index is larger than the cap of locals, locals needs to grow
	if index > len(vm.Locals)-1 {
		vm.Locals = append(vm.Locals, FeNothing{}, FeNothing{}, FeNothing{}, FeNothing{}, FeNothing{})
	}
	vm.Locals[index] = vm.Window[reg]
}

func (vm *VM) Move(instruction Instr) {
	/*
	   MOVE dest src
	   Move the contents of dest to src in the current window.
	*/
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) -> R(%d)", b, a)
	vm.Window[a] = vm.Window[b]
}

func (vm *VM) Add(instruction Instr) {
	/*
	   ADD dest src
	   Does an add operation dest += src
	   Will panic at runtime if the type doesn't assert correctly.
	*/
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) + R(%d) -> R(%d)", a, b, a)
	vm.Window[a] = vm.Window[a].(FeNumericType).Add(vm.Window[b])
}

func (vm *VM) Sub(instruction Instr) {
	/*
	   SUB dest src
	   Does a subtraction operation dest -= src
	   Will panic at runtime if the type doesn't assert correctly.
	*/
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) - R(%d) -> R(%d)", a, b, a)
	vm.Window[a] = vm.Window[a].(FeNumericType).Sub(vm.Window[b])

}

func (vm *VM) Div(instruction Instr) {
	/*
	   DIV dest src
	   Does a division operation dest /= src
	   Will panic at runtime if the type doesn't assert correctly.
	*/
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) / R(%d) -> R(%d)", a, b, a)
	vm.Window[a] = vm.Window[a].(FeNumericType).Div(vm.Window[b])
}

func (vm *VM) Mul(instruction Instr) {
	/*
	   MUL dest src
	   Does a multiplication operation dest *= src
	   Will panic at runtime if the type doesn't assert correctly.
	*/
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) * R(%d) -> R(%d)", a, b, a)
	vm.Window[a] = vm.Window[a].(FeNumericType).Mul(vm.Window[b])
}

func (vm *VM) Mod(instruction Instr) {
	/*
	   MOD dest src
	   Does a mod operation dest %= src
	   Will panic at runtime if the type doesn't assert correctly.
	*/
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) % R(%d) -> R(%d)", a, b, a)
	vm.Window[a] = vm.Window[a].(FeNumericType).Mod(vm.Window[b])
}

func (vm *VM) Cmp(instruction Instr) {
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) cmp R(%d) -> R(%d)", a, b, a)
	vm.Cf = vm.Window[a].Cmp(vm.Window[b])
}

func (vm *VM) JmpCmp(instruction Instr, cmp int) {
	a := instruction.Args[0]
	log.Debug("R(%d) == %s", a, debugCmp(cmp))
	if (vm.Cf & cmp) == cmp {
		log.Debug("True")
		vm.Pc = int(vm.Program.Consts[a].(FeInt))
		return
	}
	log.Debug("False")
}

func (vm *VM) Inc(instruction Instr) {
	a := instruction.Args[0]
	log.Debug("R(%d) + 1 -> R(%d)", a, a)
	vm.Window[a] = vm.Window[a].(FeNumericType).Add(FeInt(1))
}

func (vm *VM) Dec(instruction Instr) {
	a := instruction.Args[0]
	log.Debug("R(%d) - 1 -> R(%d)", a, a)
	vm.Window[a] = vm.Window[a].(FeNumericType).Sub(FeInt(1))
}

func (vm *VM) NewList(instruction Instr) {
	a := instruction.Args[0]
	log.Debug("[0] -> R(%d)", a)
	vm.Window[a] = make(FeList, 10)
}

func (vm *VM) AddList(instruction Instr) {
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("R(%d) push R(%d) -> R(%d)")
	vm.Window[a].(FeListType).Push(vm.Window[b])
}

func (vm *VM) PopList(instruction Instr) {
	a := instruction.Args[0]
	b := instruction.Args[1]
	log.Debug("pop R(%d) -> R(%d)", b, a)
	vm.Window[a] = vm.Window[b].(FeListType).Pop()
}

func (vm *VM) IndexList(instruction Instr) {
	a := instruction.Args[0]
	b := instruction.Args[1]
	c := instruction.Args[2]
	vm.Window[c] = vm.Window[a].(FeListType).Index(int(vm.Window[b].(FeInt)))
}

func (vm *VM) Call(instruction Instr) {
	a := instruction.Args[0]
	name := vm.Program.Consts[a].String()
	log.Debug("call %s", name)
	vm.LazyLoadFunction(name)
	// Store the return location in the 0th register
	vm.Window[8] = FeInt(vm.Pc)
	vm.Pc = vm.JumpTable[name]
	vm.IncrWindow()
}

func (vm *VM) Sys(instruction Instr) {
	log.Debug("Making System Call")
	cval := instruction.Args[0]
	name := vm.Program.Consts[cval].String()
	log.Debug("Calling %s", name)
	vm.System.Invoke(name, vm.Window)
}

func (vm *VM) Ret(instruction Instr) {
	log.Debug("ret")
	vm.DecrWindow()
	whereToJmpBackTo := vm.Window[8].(FeInt)
	vm.Pc = int(whereToJmpBackTo)
}
