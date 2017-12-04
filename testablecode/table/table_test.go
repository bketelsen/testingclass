package table

import (
	"fmt"
	"testing"
)

type TestCase struct {
	A      int
	B      int
	Result int
}

func TestMultiply(t *testing.T) {
	tc := TestCase{A: 1, B: 2, Result: 2}
	res := Multiply(tc.A, tc.B)
	if res != tc.Result {
		t.Errorf("Expected %d, got %d", tc.Result, res)
	}
}

var tests = []TestCase{
	{A: 1, B: 2, Result: 2},
	{A: 10, B: 10, Result: 100},
}

func TestTableMultiply(t *testing.T) {
	for _, tc := range tests {
		res := Multiply(tc.A, tc.B)
		if res != tc.Result {
			t.Errorf("Expected %d, got %d", tc.Result, res)
		}
	}
}

var anontests = []struct {
	A      int
	B      int
	Result int
}{
	{A: 1, B: 2, Result: 2},
	{A: 10, B: 10, Result: 100},
}

func TestAnonMultiply(t *testing.T) {
	for _, tc := range anontests {
		res := Multiply(tc.A, tc.B)
		if res != tc.Result {
			t.Errorf("Expected %d, got %d", tc.Result, res)
		}
	}
}
func TestSubMultiply(t *testing.T) {
	for _, tc := range anontests {
		t.Run(fmt.Sprintf("%d times %d", tc.A, tc.B), func(t *testing.T) {
			res := Multiply(tc.A, tc.B)
			if res != tc.Result {
				t.Errorf("Expected %d, got %d", tc.Result, res)
			}
		})
	}
}
