package types

import (
	test "testing"
)

func TestBoolString(t *test.T) {
	a := FeBool(true)
	if a.String() != "true" {
		t.Error("'true' != 'true'")
	}
}

func TestBoolEqual(t *test.T) {
	a := FeBool(true)
	b := FeBool(true)
	if (a.Cmp(b) & Equal) != Equal {
		t.Errorf("%s is not equal to %s", a, b)
	}
}

func TestBoolNotEqual(t *test.T) {
	a := FeBool(false)
	b := FeBool(true)
	if (a.Cmp(b) & NotEqual) != NotEqual {
		t.Errorf("%s is equal to %s", a, b)
	}
}
