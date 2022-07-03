package slice

import "litu/option"

type MapFn[T, U any] func(*T) U

type PredFn[T any] func(*T) bool

func Map[T, U any](a []T, f MapFn[T, U]) []U {
	b := make([]U, len(a))
	for i := range a {
		b[i] = f(&a[i])
	}
	return b
}

func MapInPlace[T any](a []T, f MapFn[T, T]) []T {
	for i := range a {
		a[i] = f(&a[i])
	}
	return a
}

func Join[T any](a []T, sep T) []T {
	if len(a) == 0 {
		return []T{}
	}
	t := make([]T, len(a)*2-1)
	j := 0
	for i := 0; i < len(a); i++ {
		if j > 0 {
			t[j] = sep
			j++
		}
		t[j] = a[i]
		j++
	}
	return t
}

func Filter[T any](a []T, pred PredFn[T]) []T {
	b := make([]T, 0, len(a))
	for i := range a {
		if pred(&a[i]) {
			b = append(b, a[i])
		}
	}
	return b
}

func FilterMap[T, U any](a []T, f MapFn[T, option.Option[U]]) []U {
	b := make([]U, 0, len(a))
	for i := range a {
		v := f(&a[i])
		if v.IsSome() {
			b = append(b, v.Inner())
		}
	}
	return b
}

func Flat[T any](a [][]T) []T {
	n := 0
	for _, t := range a {
		n += len(t)
	}
	r := make([]T, 0, n)
	for _, t := range a {
		for _, v := range t {
			r = append(r, v)
		}
	}
	return r
}

func FlatMap[T, U any](a [][]T, f MapFn[T, U]) []U {
	n := 0
	for _, t := range a {
		n += len(t)
	}
	r := make([]U, 0, n)
	for _, t := range a {
		for i := range t {
			r = append(r, f(&t[i]))
		}
	}
	return r
}
