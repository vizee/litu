package assert

import "testing"

func TestAssertEq(t *testing.T) {
	a := 1
	b := 2
	AssertEq(a, b)
}
