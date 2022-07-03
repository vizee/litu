package assert

import (
	"fmt"
	"strings"
)

func AssertEq[T comparable](lhs T, rhs T) {
	if assertEnabled {
		if lhs != rhs {
			panic(fmt.Sprintf("left(%v) is not equal to right(%v)", lhs, rhs))
		}
	}
}

func Assert(eq func() bool, msg ...string) {
	if assertEnabled {
		if eq() {
			panic(strings.Join(msg, " "))
		}
	}
}
