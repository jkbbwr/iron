package types

type FeList []FeType

func (this FeList) String() string {
	return this.String()
}

func (this FeList) Cmp(that FeType) int {
	return NotEqual
}

func (this *FeList) Push(other FeType) {
	(*this) = append(*this, other)
}

func (this *FeList) Pop() FeType {
	out := (*this)[len(*this)-1]
	(*this) = (*this)[:len(*this)-1]
	return out
}

func (this *FeList) Index(index int) FeType {
	out := (*this)[index]
	return out
}
