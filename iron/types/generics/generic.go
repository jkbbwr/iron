package generics

import . "iron/types"

/*
These types should NOT be used at runtime. They are simply here for faster prototyping of the interfaces.
When you wish to add a type to github.com.jkbbwr.iron, copy the interface you want to implement, and this generic interface
Then replace T with the type name prefixed with Fe and replace struct{} with whatever type you are supporting.
*/

type T struct{}

func (this T) String() string {
	panic("T: Not implemented")
}

func (this T) Cmp(that FeType) int {
	/*
		// Two different types are never equal
		other, ok := that.(T)
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
	*/
	panic("T: Not implemented")
}
