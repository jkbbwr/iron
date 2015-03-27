package sys

/*
func (vm *VM) Open(instruction Instr) {
    dest := instruction.Args[0]
    path := vm.Window[instruction.Args[1]].(FeString).String()
    mode := vm.Window[instruction.Args[2]].(FeString).String()
    log.Debug("Creating FeFile{%s, %s}", path, mode)

    frags := strings.SplitN(path, "+", 2)
    if len(frags) != 2 {
        log.Fatalf("Given path has incorrect format. Should be type+path got %s", path)
    }

    ftype := frags[0]
    path = frags[1]

    file := NewFile(ftype, path, mode)
    vm.Window[dest] = file
}

func (vm *VM) Read(instruction Instr) {
    dest := instruction.Args[0]
    file := vm.Window[instruction.Args[1]].(FeFile)
    size := vm.Window[instruction.Args[2]].(FeInt)
    vm.Window[dest] = FeString(file.Read(int(size)))
}*/
