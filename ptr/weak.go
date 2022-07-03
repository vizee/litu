package ptr

import "unsafe"

type Weak[T any] uintptr

func (p Weak[T]) Ptr() *T {
	return (*T)(unsafe.Pointer(p))
}

func (p Weak[T]) Deref() T {
	return *p.Ptr()
}

func (p Weak[T]) Write(v T) {
	*p.Ptr() = v
}

func (p Weak[T]) Copy(src Weak[T]) {
	*p.Ptr() = *src.Ptr()
}

func (p Weak[T]) Swap(x Weak[T]) {
	t := *p.Ptr()
	*p.Ptr() = *x.Ptr()
	*x.Ptr() = t
}

func (p Weak[T]) Offset(n int) Weak[T] {
	return p + Weak[T](unsafe.Sizeof(*(*T)(nil))*uintptr(n))
}

func (p Weak[T]) Width(b Weak[T]) int {
	return int(uintptr(p-b)) / int(unsafe.Sizeof(*(*T)(nil)))
}
