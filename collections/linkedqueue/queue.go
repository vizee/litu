package linkedqueue

type qlink[T any] struct {
	v    T
	link [2]*qlink[T]
}

type Queue[T any] struct {
	link [2]*qlink[T]
}

func (q *Queue[T]) Empty() bool {
	return q.link[0] == nil
}

func (q *Queue[T]) push(v T, d int) {
	d &= 1
	t := (d + 1) & 1

	e := &qlink[T]{
		v: v,
	}
	e.link[d] = q.link[d]
	if e.link[d] != nil {
		e.link[d].link[t] = e
	}

	q.link[d] = e
	if q.link[t] == nil {
		q.link[t] = e
	}
}

func (q *Queue[T]) pop(d int) (T, bool) {
	d &= 1
	t := (d + 1) & 1

	e := q.link[d]
	if e == nil {
		var zero T
		return zero, false
	}

	q.link[d] = e.link[d]
	if q.link[d] != nil {
		q.link[d].link[t] = nil
	} else {
		q.link[t] = nil
	}
	return e.v, true
}

func (q *Queue[T]) PushFront(v T) {
	q.push(v, 0)
}

func (q *Queue[T]) PopFront() (T, bool) {
	return q.pop(0)
}

func (q *Queue[T]) PushBack(v T) {
	q.push(v, 1)
}

func (q *Queue[T]) PopBack() (T, bool) {
	return q.pop(1)
}

func (q *Queue[T]) Collect() []T {
	n := 0
	for e := q.link[0]; e != nil; e = e.link[0] {
		n++
	}
	l := make([]T, n)
	e := q.link[0]
	for i := range l {
		l[i] = e.v
		e = e.link[0]
	}
	return l
}
