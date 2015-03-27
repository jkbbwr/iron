package types

type FeString string

func (this FeString) String() string {
	return string(this)
}

func (this FeString) Cmp(that FeType) int {

	// Two different types are never equal
	other, ok := that.(FeString)
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
