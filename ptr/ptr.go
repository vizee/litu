package ptr

import "unsafe"

func Read[T any](p unsafe.Pointer) T {
	return *(*T)(p)
}

type sliceView struct {
	p unsafe.Pointer
	n int
	m int
}

func AsSlice[T any](p unsafe.Pointer, n int) []T {
	s := sliceView{
		p: p,
		n: n,
		m: n,
	}
	return *(*[]T)(unsafe.Pointer(&s))
}

func FromSlice[T any](s []T) *T {
	return *(**T)(unsafe.Pointer(&s))
}

func Write[T any](p unsafe.Pointer, v T) {
	*(*T)(p) = v
}

func Copy[T any](dst *T, src *T) {
	*dst = *src
}

func Swap[T any](a *T, b *T) {
	t := *a
	*a = *b
	*b = t
}

func Less[T any](lhs *T, rhs *T) bool {
	return uintptr(unsafe.Pointer(lhs)) < uintptr(unsafe.Pointer(rhs))
}

func Offset[T any](p *T, n int) *T {
	return (*T)(unsafe.Add(unsafe.Pointer(p), int(unsafe.Sizeof(*(*T)(nil)))*n))
}

func Width[T any](p *T, b *T) int {
	return int(uintptr(unsafe.Pointer(p))-uintptr(unsafe.Pointer(b))) / int(unsafe.Sizeof(*(*T)(nil)))
}
