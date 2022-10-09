package slice

import (
	"github.com/vizee/litu/cmp"
	"github.com/vizee/litu/constraints"
	"github.com/vizee/litu/option"
)

func Reduce[T, U any](a []T, v U, f func(U, *T) U) U {
	for i := range a {
		v = f(v, &a[i])
	}
	return v
}

func Accumulate[T constraints.Accumulate](a []T) T {
	var r T
	for i := range a {
		r += a[i]
	}
	return r
}

func Max[T constraints.Number](a []T) option.Option[T] {
	if len(a) > 0 {
		r := a[0]
		for i := 1; i < len(a); i++ {
			if a[i] > r {
				r = a[i]
			}
		}
		return option.Some(r)
	} else {
		return option.None[T]()
	}
}

func Min[T constraints.Number](a []T) option.Option[T] {
	if len(a) > 0 {
		r := a[0]
		for i := 1; i < len(a); i++ {
			if a[i] < r {
				r = a[i]
			}
		}
		return option.Some(r)
	} else {
		return option.None[T]()
	}
}

func MaxBy[T any](a []T, cmp cmp.Cmp[*T]) option.Option[T] {
	if len(a) > 0 {
		r := a[0]
		for i := 1; i < len(a); i++ {
			if cmp(&a[i], &r) > 0 {
				r = a[i]
			}
		}
		return option.Some(r)
	} else {
		return option.None[T]()
	}
}

func MinBy[T any](a []T, cmp cmp.Cmp[*T]) option.Option[T] {
	if len(a) > 0 {
		r := a[0]
		for i := 1; i < len(a); i++ {
			if cmp(&a[i], &r) < 0 {
				r = a[i]
			}
		}
		return option.Some(r)
	} else {
		return option.None[T]()
	}
}

func CountBy[T any](a []T, pred PredFn[T]) int {
	n := 0
	for i := range a {
		if pred(&a[i]) {
			n++
		}
	}
	return n
}

func CountPtrBy[T any](a []*T, pred PredFn[T]) int {
	n := 0
	for i := range a {
		if pred(a[i]) {
			n++
		}
	}
	return n
}

func GroupBy[T any, K comparable](a []T, key MapFn[T, K]) map[K][]T {
	m := make(map[K][]T)
	for i := range a {
		k := key(&a[i])
		m[k] = append(m[k], a[i])
	}
	return m
}

func GroupSortedBy[T any](a []T, eq cmp.Equal[*T]) [][]T {
	var (
		r [][]T
		g []T
	)

	if len(a) > 0 {
		g = []T{a[0]}
	}
	i := 0
	for j := 1; j < len(a); j++ {
		if eq(&a[i], &a[j]) {
			g = append(g, a[j])
			continue
		}
		r = append(r, g)
		g = []T{a[j]}
		i = j
	}

	if len(g) > 0 {
		r = append(r, g)
	}

	return r
}
