package types

import "fmt"

type FeBool bool

func (this FeBool) String() string {
	return fmt.Sprintf("%t", this)
}

func (this FeBool) Cmp(that FeType) int {
	// Two different types are never equal
	other, ok := that.(FeBool)
	if !ok {
		return NotEqual
	}
	var flags int
	if this == other {
		flags |= Equal
	}
	if this != other {
		flags |= NotEqual
	}
	return flags

}
