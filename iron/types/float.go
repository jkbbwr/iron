package types

type FeFloat float64

func (this FeFloat) String() string {
	panic("FeFloat: Not implemented")
}

func (this FeFloat) Add(that FeType) FeType {
	return this + that.(FeFloat)
}

func (this FeFloat) Sub(that FeType) FeType {
	return this - that.(FeFloat)
}

func (this FeFloat) Mul(that FeType) FeType {
	return this * that.(FeFloat)
}

func (this FeFloat) Div(that FeType) FeType {
	return this / that.(FeFloat)
}

func (this FeFloat) Cmp(that FeType) int {
	// Two different types are never equal
	other, ok := that.(FeFloat)
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
