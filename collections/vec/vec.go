package vec

import "github.com/vizee/litu/option"

type Vec[T any] struct {
	a []T
}

func (v *Vec[T]) Slice() []T {
	return v.a
}

func (v *Vec[T]) Cap() int {
	return cap(v.a)
}

func (v *Vec[T]) Len() int {
	return len(v.a)
}

func (v *Vec[T]) Clear() {
	v.a = v.a[:0]
}

func (v *Vec[T]) Get(idx int) option.Option[T] {
	if idx < len(v.a) {
		return option.Some(v.a[idx])
	} else {
		return option.None[T]()
	}
}

func (v *Vec[T]) GetRef(idx int) *T {
	if idx < len(v.a) {
		return &v.a[idx]
	} else {
		return nil
	}
}

func (v *Vec[T]) Push(x T) {
	v.a = append(v.a, x)
}

func (v *Vec[T]) Pop() option.Option[T] {
	if len(v.a) > 0 {
		x := v.a[len(v.a)-1]
		v.a = v.a[:len(v.a)-1]
		return option.Some(x)
	} else {
		return option.None[T]()
	}
}

func (v *Vec[T]) Append(x []T) {
	v.a = append(v.a, x...)
}

func (v *Vec[T]) Reserve(additional int) {
	if cap(v.a)-len(v.a) >= additional {
		return
	}
	a := make([]T, len(v.a), len(v.a)+additional)
	copy(a, v.a)
	v.a = a
}

func (v *Vec[T]) Resize(additional int) {
	v.Reserve(additional)
	v.a = v.a[0 : len(v.a)+additional : cap(v.a)]
	// TODO reset a[len(a):len(a)+additional]
}

func (v *Vec[T]) Set(idx int, x T) {
	if idx >= len(v.a) {
		v.Resize(idx - len(v.a) + 1)
	}
	v.a[idx] = x
}

func (v *Vec[T]) Insert(idx int, x T) {
	if idx < len(v.a) {
		v.Resize(1)
		copy(v.a[idx+1:], v.a[idx:])
	} else {
		v.Resize(idx - len(v.a) + 1)
	}
	v.a[idx] = x
}

func (v *Vec[T]) Remove(idx int) option.Option[T] {
	if idx < len(v.a) {
		x := v.a[idx]
		copy(v.a[idx:], v.a[idx+1:])
		v.a = v.a[:len(v.a)-1]
		return option.Some(x)
	} else {
		return option.None[T]()
	}
}

func (v *Vec[T]) SwapRemove(idx int) option.Option[T] {
	if idx < len(v.a) {
		x := v.a[idx]
		v.a[idx] = v.a[len(v.a)-1]
		v.a = v.a[:len(v.a)-1]
		return option.Some(x)
	} else {
		return option.None[T]()
	}
}

func FromSlice[T any](a []T) Vec[T] {
	return Vec[T]{a: a}
}
