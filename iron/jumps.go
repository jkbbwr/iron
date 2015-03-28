package iron

import "github.com/op/go-logging"

var log = logging.MustGetLogger("FeVM")

func (vm *VM) LazyLoadFunction(name string) {
	log.Debug("Loading function %s", name)
	// Load the function
	function := vm.Program.Funcs[name]
	vm.JumpTable[name] = vm.Program.CodeLength
	log.Debug("Extending code length from %d", vm.Program.CodeLength)
	vm.Program.Instructions = append(vm.Program.Instructions, function.Instructions...)
	vm.Program.CodeLength += len(function.Instructions)
	log.Debug("Code length %d", vm.Program.CodeLength)
}

func debugJmp(vm *VM) {
	log.Debug("Jump Table:")
	for k, v := range vm.JumpTable {
		log.Debug("    Function %s at pc %d", k, v)
	}
}
