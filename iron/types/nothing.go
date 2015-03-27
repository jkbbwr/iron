package types

type FeNothing struct{}

func (this FeNothing) String() string {
	return "FeNothing"
}

func (this FeNothing) Cmp(that FeType) int {
	// Two different types are never equal
	_, ok := that.(FeNothing)
	if !ok {
		return NotEqual
	}
	return Equal
}
