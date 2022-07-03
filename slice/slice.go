package slice

import "litu/option"

func Get[T any](a []T, idx int) option.Option[T] {
	if idx < len(a) {
		return option.Some(a[idx])
	} else {
		return option.None[T]()
	}
}

func Swap[T any](a []T, i int, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func Clone[T any](a []T) []T {
	r := make([]T, len(a))
	copy(r, a)
	return r
}

func Reverse[T any](a []T) {
	l := 0
	r := len(a) - 1
	for l < r {
		Swap(a, l, r)
		l++
		r--
	}
}

func Shrink[T any](a []T) []T {
	if cap(a) > len(a) {
		return Clone(a)
	} else {
		return a
	}
}

func Repeat[T any](n int, v T) []T {
	r := make([]T, n)
	for i := 0; i < len(r); i++ {
		r[i] = v
	}
	return r
}

func RepeatSlice[T any](n int, v ...T) []T {
	r := make([]T, 0, n*len(v))
	for i := 0; i < n; i++ {
		r = append(r, v...)
	}
	return r
}

func ForEach[T any](a []T, f func(*T)) {
	for i := range a {
		f(&a[i])
	}
}
