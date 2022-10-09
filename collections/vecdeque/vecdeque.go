package vecdeque

import "github.com/vizee/litu/option"

type VecDeque[T any] struct {
	head uint
	tail uint
	buf  []T
}

func (q *VecDeque[T]) wrapIdx(idx uint) uint {
	return idx % uint(len(q.buf))
}

func (q *VecDeque[T]) resizeCap(n int) {
	buf := make([]T, n)
	m := q.Len()
	i := q.wrapIdx(q.head)
	j := q.wrapIdx(q.tail)
	if i <= j {
		copy(buf, q.buf[i:j])
	} else {
		mid := copy(buf, q.buf[i:])
		copy(buf[mid:], q.buf[:j])
	}
	q.head = 0
	q.tail = uint(m)
	q.buf = buf
}

func (q *VecDeque[T]) Cap() int {
	return len(q.buf)
}

func (q *VecDeque[T]) Len() int {
	return int(q.tail - q.head)
}

func (q *VecDeque[T]) Empty() bool {
	return q.head == q.tail
}

func (q *VecDeque[T]) Full() bool {
	return q.tail-q.head == uint(len(q.buf))
}

func (q *VecDeque[T]) Reset(capacity int) {
	q.head = 0
	q.tail = 0
	q.buf = make([]T, capacity)
}

func (q *VecDeque[T]) Clear() {
	q.head = 0
	q.tail = 0
}

func (q *VecDeque[T]) Reserve(additional int) {
	n := q.Len()
	rem := q.Cap() - n
	if rem >= additional {
		return
	}
	q.resizeCap(n + additional)
}

func (q *VecDeque[T]) growCap() {
	n := len(q.buf)
	if n < 1024 {
		q.resizeCap(n + n)
	} else {
		q.resizeCap(n + n/4)
	}
}

func (q *VecDeque[T]) PushFront(v T) {
	if q.Full() {
		q.growCap()
	}
	q.head--
	q.buf[q.wrapIdx(q.head)] = v
}

func (q *VecDeque[T]) PopFront() option.Option[T] {
	if !q.Empty() {
		v := q.buf[q.wrapIdx(q.head)]
		q.head++
		return option.Some(v)
	} else {
		return option.None[T]()
	}
}

func (q *VecDeque[T]) PushBack(v T) {
	if q.Full() {
		q.growCap()
	}
	q.buf[q.wrapIdx(q.tail)] = v
	q.tail++
}

func (q *VecDeque[T]) PopBack() option.Option[T] {
	if !q.Empty() {
		q.tail++
		return option.Some(q.buf[q.wrapIdx(q.tail)])
	} else {
		return option.Option[T]{}
	}
}

func (q *VecDeque[T]) First() option.Option[T] {
	if !q.Empty() {
		return option.Some(q.buf[q.wrapIdx(q.head)])
	} else {
		return option.None[T]()
	}
}

func (q *VecDeque[T]) FirstRef() *T {
	if !q.Empty() {
		return &q.buf[q.wrapIdx(q.head)]
	} else {
		return nil
	}
}

func (q *VecDeque[T]) Last() option.Option[T] {
	if !q.Empty() {
		return option.Some(q.buf[q.wrapIdx(q.tail-1)])
	} else {
		return option.None[T]()
	}
}

func (q *VecDeque[T]) LastRef() *T {
	if !q.Empty() {
		return &q.buf[q.wrapIdx(q.tail-1)]
	} else {
		return nil
	}
}

func (q *VecDeque[T]) Get(idx int) option.Option[T] {
	if q.tail-q.head > uint(idx) {
		return option.Some(q.buf[q.wrapIdx(q.head+uint(idx))])
	} else {
		return option.None[T]()
	}
}

func New[T any](capacity int) *VecDeque[T] {
	return &VecDeque[T]{
		head: 0,
		tail: 0,
		buf:  make([]T, capacity),
	}
}
