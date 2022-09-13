package option

type Option[T any] struct {
	v    T
	some bool
}

func (o *Option[T]) Inner() T {
	return o.v
}

func (o *Option[T]) IsSome() bool {
	return o.some
}

func (o *Option[T]) AsRef() Ref[T] {
	if o.some {
		return Ref[T]{ptr: &o.v}
	} else {
		return Ref[T]{ptr: nil}
	}
}

func (o *Option[T]) ToSlice() []T {
	if o.some {
		return []T{o.v}
	} else {
		return nil
	}
}

func (o *Option[T]) Take() Option[T] {
	if o.some {
		r := *o
		*o = Option[T]{}
		return r
	} else {
		return *o
	}
}

func (o *Option[T]) OrElse(f func() Option[T]) Option[T] {
	if !o.some {
		return f()
	} else {
		return *o
	}
}

func (o *Option[T]) Into() (T, bool) {
	return o.v, o.some
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		v:    v,
		some: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func Map[T, U any](o *Option[T], f func(*T) U) Option[U] {
	if o.some {
		return Some(f(&o.v))
	} else {
		return None[U]()
	}
}

func AndThen[T, U any](o *Option[T], f func(*T) Option[U]) Option[U] {
	if o.some {
		return f(&o.v)
	} else {
		return None[U]()
	}
}
