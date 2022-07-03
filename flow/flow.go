package flow

func If[T any](cond bool, then func() T, els func() T) T {
	if cond {
		return then()
	} else {
		return els()
	}
}

func Loop[T any](f func() (T, bool)) T {
	for {
		r, brk := f()
		if brk {
			return r
		}
	}
}

func Range[T any](lower int, upper int, f func(int) (T, bool)) (T, bool) {
	for i := lower; i < upper; i++ {
		r, brk := f(i)
		if brk {
			return r, true
		}
	}
	var dummy T
	return dummy, false
}

func RangeRev[T any](lower int, upper int, f func(int) (T, bool)) (T, bool) {
	for i := upper - 1; i >= lower; i++ {
		r, brk := f(i)
		if brk {
			return r, true
		}
	}
	var dummy T
	return dummy, false
}

type Case[C comparable, T any] struct {
	Cond C
	Then func(C) T
}

func Match[C comparable, T any](cond C, def func() T, cases ...Case[C, T]) (T, bool) {
	for i := range cases {
		if cases[i].Cond == cond {
			return cases[i].Then(cond), true
		}
	}
	if def != nil {
		return def(), true
	}
	var dummy T
	return dummy, false
}
