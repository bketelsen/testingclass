package scope

import "testing"

func TestMultiply(t *testing.T) {
	m := new(IntCalculator)
	m.Multiply(5, 6)
	if m.Result != 30 {
		t.Errorf("got %d, expected %d", m.Result, 30)
	}
}

// adding a helper allows you to keep state contained
// and reduce boilerplate code
func multiply(a, b int) *IntCalculator {
	m := new(IntCalculator)
	m.Multiply(a, b)
	return m
}

// This code doesn't test whether 5*6 = 30
// it tests whether the result is properly stored in the
// IntCalculator type.
func TestMultiplyWithHelper(t *testing.T) {
	m := multiply(5, 6)
	if m.Result != 30 {
		t.Errorf("got %d, expected %d", m.Result, 30)
	}
}

func TestClear(t *testing.T) {
	m := multiply(5, 6)
	m.Clear()
	if m.Result != 0 {
		t.Errorf("got %d, expected %d", m.Result, 0)
	}
}
