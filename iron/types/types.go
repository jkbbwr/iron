package types

const (
	// (v & Equal) == Equal
	Equal = 1 << iota
	NotEqual
	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
)

/*
This is the base interface of every type in the language. All types *MUST* implement String
*/
type FeType interface {
	String() string
	Cmp(other FeType) int
}

type FeNumericType interface {
	Add(other FeType) FeType
	Sub(other FeType) FeType
	Mul(other FeType) FeType
	Div(other FeType) FeType
	Mod(other FeType) FeType
}

type FeListType interface {
	Push(other FeType)
	Pop() FeType
	Index(index int) FeType
}
