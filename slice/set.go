package slice

import (
	"github.com/vizee/litu/cmp"
	"github.com/vizee/litu/constraints"
)

func Union[T comparable](a []T, b []T) []T {
	if len(a) < len(b) {
		a, b = b, a
	}
	r := make([]T, 0, len(a)+len(b))
	r = append(r, a...)
	for i := range b {
		if Contains(a, &b[i]) {
			continue
		}
		r = append(r, b[i])
	}
	return r
}

func Intersect[T comparable](a []T, b []T) []T {
	if len(a) < len(b) {
		a, b = b, a
	}
	r := make([]T, 0, len(b))
	for i := range b {
		if Contains(a, &b[i]) {
			r = append(r, b[i])
		}
	}
	return r
}

func Diff[T comparable](a []T, b []T) []T {
	r := make([]T, 0, len(a))
	for i := range a {
		if !Contains(b, &a[i]) {
			r = append(r, a[i])
		}
	}
	return r
}

func UnionSorted[T constraints.Compare](a []T, b []T) []T {
	r := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else if a[i] > b[j] {
			r = append(r, b[j])
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	if i < len(a) {
		r = append(r, a[:i]...)
	}
	if j < len(b) {
		r = append(r, b[:j]...)
	}
	return r
}

func IntersectSorted[T constraints.Compare](a []T, b []T) []T {
	r := make([]T, 0, cmp.Min(len(a), len(b)))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func DiffSorted[T constraints.Compare](a []T, b []T) []T {
	r := make([]T, 0, len(a))
	for i := range a {
		if !BinarySearch(b, a[i]).IsSome() {
			r = append(r, a[i])
		}
	}
	return r
}

func DedupSorted[T comparable](a []T) []T {
	i := 0
	for j := 1; j < len(a); j++ {
		if a[i] != a[j] {
			i++
			if i != j {
				a[i] = a[j]
			}
		}
	}
	if i+1 < len(a) {
		return a[:i+1]
	} else {
		return a
	}
}
