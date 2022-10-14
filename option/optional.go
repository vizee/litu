package option

type Optional[T any] interface {
	Inner() T
	IsSome() bool
}

func Unwrap[T any, O Optional[T]](o O) T {
	if o.IsSome() {
		return o.Inner()
	}
	panic(`unwrap a none value`)
}

func UnwrapOrElse[T any, O Optional[T]](o O, orElse func() T) T {
	if o.IsSome() {
		return o.Inner()
	} else {
		return orElse()
	}
}

func ToSlice[T any, O Optional[T]](o O) []T {
	if o.IsSome() {
		return []T{o.Inner()}
	} else {
		return nil
	}
}

func Iff[T any, O Optional[T], U any](o O, then func(T) U, els ...func() U) U {
	if o.IsSome() {
		return then(o.Inner())
	} else if len(els) > 0 {
		return els[0]()
	} else {
		var zero U
		return zero
	}
}

func Map[T, U any, O Optional[T]](o O, f func(T) U) Option[U] {
	if o.IsSome() {
		return Some(f(o.Inner()))
	} else {
		return None[U]()
	}
}

func AndThen[T, U any, O Optional[T]](o O, f func(T) Option[U]) Option[U] {
	if o.IsSome() {
		return f(o.Inner())
	} else {
		return None[U]()
	}
}
