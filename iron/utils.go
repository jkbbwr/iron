package iron

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
    "github.com/jkbbwr/iron/iron/types"
)

const debugHelp = `
(s)tep - Execute the next instruction and return to debugger.
(h)elp - Get this help message

(r)egister [n] - Print the contents of register [n] in the current window
(c)onst [n] - Print the const stored at [n]
(l)ocal [n] - Print the local stored at [n]

(w)indow (p)rev - Print the previous window
(w)indow (c)urrent - Print the current window
(w)indow (n)ext - Print the next window

sys - Show system registered functions
jmp - Debug the jump table

pc - Print the PC
wp - Print the wp
cf - Print the cf
`

var opmap = [...]string{
	"LConst", "LLocal", "SLocal",
	"Move", "Clr", "Nop",
	"Jmp", "Cmp", "Add", "Sub", "Div",
	"Mul", "Inc", "Dec", "Mod", "NewList",
	"AddList", "Poplist", "Call", "Ret",
	"JmpEq", "JmpNEq", "JmpGt", "JmpGtEq",
	"JmpLt", "JmpLtEq", "Sys",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil", "nil",
	"Halt",
}

func (vm *VM) Debugger() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">>> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.Trim(line, "\n")
		command, arg, err := twoStrings(strings.SplitN(line, " ", 2))
		switch command {
		case "h":
			fallthrough
		case "help":
			log.Debug(debugHelp)
			continue
		case "s":
			fallthrough
		case "step":
			return
		case "w":
			fallthrough
		case "window":
			switch arg {
			case "n":
				fallthrough
			case "next":
				dumpNextWindow(vm)
				continue
			case "p":
				fallthrough
			case "prev":
				dumpPrevWindow(vm)
				continue
			case "c":
				fallthrough
			case "current":
				dumpWindow(vm)
				continue
			}
		case "r":
			fallthrough
		case "register":
			reg, err := strconv.ParseInt(arg, 10, 32)
			if err != nil {
				log.Warning("%s", err)
			} else {
				log.Debug("R(%d) = %+#v", reg, vm.Window[reg])
				continue
			}

		case "c":
			fallthrough
		case "const":
			n, err := strconv.ParseInt(arg, 10, 32)
			if err != nil {
				log.Warning("%s", err)
			} else {
				log.Debug("Const(%d) = %+#v", n, vm.Program.Consts[n])
				continue
			}
		case "sys":
			debugSys(vm)
			continue
		case "jmp":
			debugJmp(vm)
			continue
		case "pc":
			log.Debug("Pc = %d", vm.Pc)
			continue
		case "wp":
			log.Debug("Wp = %d", vm.Wp)
			continue
		case "cf":
			log.Debug("Cf = %s", debugCmp(vm.Cf))
			continue
		case "l":
			fallthrough
		case "local":
			n, err := strconv.ParseInt(arg, 10, 32)
			if err != nil {
				log.Warning("%s", err)
			} else {
				log.Debug("Local(%d) = %+#v", n, vm.Locals[n])
				continue
			}

		}
		log.Warning("Debug command not known.")
	}
}

func twoStrings(frags []string) (string, string, error) {
	if len(frags) == 1 {
		return frags[0], "", nil
	}
	if len(frags) != 2 {
		return "", "", errors.New("Not enough fragments")
	}
	return frags[0], frags[1], nil
}

func debugSys(vm *VM) {
    log.Debug("System Table: ")
    for k, v := range(vm.System.Table) {
        log.Debug("    Func %s: %+#v", k, v)
    }
}

func dumpWindow(vm *VM) {
	log.Debug("Window (wp=%d): %+#v", vm.Wp, vm.Window)
}

func dumpNextWindow(vm *VM) {
	wp := vm.Wp + 17
	window := vm.Register[wp : wp+25]
	log.Debug("Window (wp=%d): %+#v", wp, window)
}

func dumpPrevWindow(vm *VM) {
	if vm.Wp == 0 {
		log.Warning("I cannot back one window")
		return
	}
	wp := vm.Wp - 17
	window := vm.Register[wp : wp+25]
	log.Debug("Window (wp=%d): %+#v", wp, window)
}

func debugCmp(cmp int) string {
	switch cmp {
	case types.Equal:
		return "Equal"
	case types.NotEqual:
		return "NotEqual"
	case types.LessThan:
		return "LessThan"
	case types.LessThanOrEqual:
		return "LessThanOrEqual"
	case types.GreaterThan:
		return "GreaterThan"
	case types.GreaterThanOrEqual:
		return "GreaterThanOrEqual"
	}
	return "Nothing"
}

func debugOpcode(opcode int) string {
	return opmap[opcode]
}
