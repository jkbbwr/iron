package types

import (
	test "testing"
)

func TestNothingString(t *test.T) {
	a := FeNothing{}
	if a.String() != "FeNothing" {
		t.Error("FeNothing wasn't formatted")
	}
}

func TestNothingEqual(t *test.T) {
	a := FeNothing{}
	b := FeNothing{}
	if (a.Cmp(b) & Equal) != Equal {
		t.Errorf("%s is not equal to %s", a, b)
	}
}

func TestNothingNotNothing(t *test.T) {
	a := FeBool(false)
	b := FeNothing{}
	if (a.Cmp(b)) != NotEqual {
		t.Errorf("Type assertion messed up")
	}
}
