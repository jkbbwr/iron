package types

import (
	test "testing"
)

func TestIntAdd(t *test.T) {
	a := FeInt(5)
	b := FeInt(5)
	c := (a.Add(b)).(FeInt)
	if c != 10 {
		t.Error("5 + 5 != 10")
	}
}

func TestIntSub(t *test.T) {
	a := FeInt(5)
	b := FeInt(5)
	c := (a.Sub(b)).(FeInt)
	if c != 0 {
		t.Error("5 - 5 != 0")
	}
}

func TestIntEqual(t *test.T) {
	a := FeInt(5)
	b := FeInt(5)
	if (a.Cmp(b) & Equal) != Equal {
		t.Errorf("%s is not equal to %s", a, b)
	}
}

func TestIntNotEqual(t *test.T) {
	a := FeInt(5)
	b := FeInt(6)
	if (a.Cmp(b) & NotEqual) != NotEqual {
		t.Errorf("%s is equal to %s", a, b)
	}
}

func TestIntLessThan(t *test.T) {
	a := FeInt(5)
	b := FeInt(6)
	if (a.Cmp(b) & LessThan) != LessThan {
		t.Errorf("%s is not less than %s", a, b)
	}
}

func TestIntLessThanOrEqual(t *test.T) {
	a := FeInt(5)
	b := FeInt(5)
	if (a.Cmp(b) & LessThanOrEqual) != LessThanOrEqual {
		t.Errorf("%s is not less than or equal to %s", a, b)
	}
}

func TestIntGreaterThan(t *test.T) {
	a := FeInt(6)
	b := FeInt(5)
	if (a.Cmp(b) & GreaterThan) != GreaterThan {
		t.Errorf("%s is not greater than %s", a, b)
	}
}

func TestIntGreaterThanOrEqual(t *test.T) {
	a := FeInt(5)
	b := FeInt(5)
	if (a.Cmp(b) & GreaterThanOrEqual) != GreaterThanOrEqual {
		t.Errorf("%s is not greater than or equal to %s", a, b)
	}
}
