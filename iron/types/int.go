package types

import "fmt"

type FeInt int

func (this FeInt) String() string {
	return fmt.Sprintf("%d", this)
}

func (this FeInt) Add(that FeType) FeType {
	return this + that.(FeInt)
}

func (this FeInt) Sub(that FeType) FeType {
	return this - that.(FeInt)
}

func (this FeInt) Mul(that FeType) FeType {
	return this * that.(FeInt)
}

func (this FeInt) Div(that FeType) FeType {
	return this / that.(FeInt)
}

func (this FeInt) Mod(that FeType) FeType {
	return this % that.(FeInt)
}

func (this FeInt) Push(other FeType) {
	panic("FeInt: Not implemented")
}

func (this FeInt) Pop() FeType {
	panic("FeInt: Not implemented")
}

func (this FeInt) Cmp(that FeType) int {
	// Two different types are never equal
	other, ok := that.(FeInt)
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
	if this >= other {
		flags |= GreaterThanOrEqual
	}
	if this <= other {
		flags |= LessThanOrEqual
	}
	if this < other {
		flags |= LessThan
	}
	if this > other {
		flags |= GreaterThan
	}
	return flags
}
