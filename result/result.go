package result

type R[T any] struct {
	ok  T
	err error
}

func (r *R[T]) IsOk() bool {
	return r.err == nil
}

func (r *R[T]) Ok() T {
	return r.ok
}

func (r *R[T]) Err() error {
	return r.err
}

func (r *R[T]) Unwrap() T {
	if r.IsOk() {
		return r.ok
	}
	panic(r.err)
}

func (r *R[T]) UnwrapOr(v T) T {
	if r.IsOk() {
		return r.ok
	}
	return v
}

func (r *R[T]) UnwrapOrElse(f func() T) T {
	if r.IsOk() {
		return r.ok
	}
	return f()
}

func (r *R[T]) OrElse(e func(error) R[T]) R[T] {
	if r.IsOk() {
		return *r
	} else {
		return e(r.err)
	}
}

func Map[T, U any](r R[T], f func(T) U) R[U] {
	if r.IsOk() {
		return R[U]{ok: f(r.ok)}
	}
	return R[U]{err: r.err}
}

func AndThen[T, U any](r R[T], then func(T) R[U]) R[U] {
	if r.IsOk() {
		return then(r.ok)
	} else {
		return R[U]{err: r.err}
	}
}
