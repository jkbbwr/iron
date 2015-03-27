package sys

import (
	"fmt"
    "github.com/jkbbwr/iron/iron/types"
)

func SysPrint(window []types.FeType) {
	arg := window[17]
	newLine, ok := window[18].(types.FeBool)
	if !ok {
		// The argument was the wrong type but its optional so assume newLine isn't wanted.
		newLine = false
	}
	if newLine {
		log.Debug("newline = true")
		fmt.Printf("%s\n", arg)
		return
	}
	fmt.Printf("%s", arg)
}
