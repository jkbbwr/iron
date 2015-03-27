package types

import (
	test "testing"
)

func TestStringCmpEqual(t *test.T) {
	a := FeString("test")
	b := FeString("test")
	if (a.Cmp(b) & Equal) != Equal {
		t.Errorf("%s is not equal to %s", a, b)
	}
}

func TestCmpNotEqual(t *test.T) {
	a := FeString("test")
	b := FeString("testing")
	if (a.Cmp(b) & NotEqual) != NotEqual {
		t.Errorf("%s is equal to %s", a, b)
	}
}
