package either

import (
	"github.com/vizee/litu/option"
)

type Either[L, R any] struct {
	mark  byte
	left  L
	right R
}

func (e *Either[L, R]) IsLeft() bool {
	return e.mark == 1
}

func (e *Either[L, R]) IsRight() bool {
	return e.mark == 2
}

func (e *Either[L, R]) Left() option.Option[L] {
	if e.mark == 1 {
		return option.Some(e.left)
	} else {
		return option.None[L]()
	}
}

func (e *Either[L, R]) Right() option.Option[R] {
	if e.mark == 2 {
		return option.Some(e.right)
	} else {
		return option.None[R]()
	}
}

func (e *Either[L, R]) LeftRef() *L {
	if e.mark == 1 {
		return &e.left
	} else {
		return nil
	}
}

func (e *Either[L, R]) RightRef() *R {
	if e.mark == 2 {
		return &e.right
	} else {
		return nil
	}
}

func Left[L, R any](v L) Either[L, R] {
	return Either[L, R]{
		mark: 1,
		left: v,
	}
}

func Right[L, R any](v R) Either[L, R] {
	return Either[L, R]{
		mark:  2,
		right: v,
	}
}

func Match[L, R, T any](e *Either[L, R], left func(*L) T, right func(*R) T) T {
	switch e.mark {
	case 1:
		return left(&e.left)
	case 2:
		return right(&e.right)
	default:
		panic("illegal Either value")
	}
}

func Unwrap[T any](e *Either[T, T]) T {
	switch e.mark {
	case 1:
		return e.left
	case 2:
		return e.right
	default:
		panic("illegal Either value")
	}
}
