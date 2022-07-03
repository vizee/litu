package option

type Ref[T any] struct {
	ptr *T
}

func (r Ref[T]) Inner() T {
	return *r.ptr
}

func (r Ref[T]) IsSome() bool {
	return r.ptr != nil
}

func (r *Ref[T]) Ptr() *T {
	return r.ptr
}

func RefFrom[T any](p *T) Ref[T] {
	return Ref[T]{ptr: p}
}
