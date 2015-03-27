package types

import test "testing"

func TestListPush(t *test.T) {
	a := make(FeList, 0)
	a.Push(FeInt(5))
	if len(a) != 1 {
		t.Errorf("FeList length not as expected, actally %d", len(a))
	}
}

func TestListPop(t *test.T) {
	a := make(FeList, 1)
	a[0] = FeInt(5)
	b, ok := a.Pop().(FeInt)
	if !ok {
		t.Error("Popped value failed expected type assertion")
	}
	if (b.Cmp(FeInt(5)) & Equal) != Equal {
		t.Error("Popped value wasn't equal to expected value")
	}
}

func TestListPopOrder(t *test.T) {
	a := make(FeList, 2)
	a[0] = FeInt(10)
	a[1] = FeInt(5)

	b, ok := a.Pop().(FeInt)
	if !ok {
		t.Error("Popped value failed expected type assertion")
	}
	if (b.Cmp(FeInt(5)) & Equal) != Equal {
		t.Errorf("Popped value wasn't equal to expected value, actual value was %d", int(b))
	}
	b, ok = a.Pop().(FeInt)
	if !ok {
		t.Error("Popped value failed expected type assertion")
	}
	if (b.Cmp(FeInt(10)) & Equal) != Equal {
		t.Errorf("Popped value wasn't equal to expected value, actual value was %d", int(b))
	}
}
