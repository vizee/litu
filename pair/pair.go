package pair

type Pair[A, B any] struct {
	A A
	B B
}

func (p *Pair[A, B]) Unzip() (A, B) {
	return p.A, p.B
}

func Array[T any](pair Pair[T, T]) [2]T {
	return [2]T{pair.A, pair.B}
}
