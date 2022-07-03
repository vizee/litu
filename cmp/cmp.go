package cmp

import . "github.com/vizee/litu/constraints"

type Cmp[T any] func(a, b T) int

type Equal[T any] func(a, b T) bool

func Min[T Compare](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max[T Compare](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func Clamp[T Compare](min, max, v T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func Less[T Compare](a *T, b *T) bool {
	return *a < *b
}

func Greater[T Compare](a *T, b *T) bool {
	return *a > *b
}
