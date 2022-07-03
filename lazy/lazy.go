package lazy

import (
	"sync"
	"sync/atomic"
)

type Lazy[T any] struct {
	state int32
	v     T
	init  func() T
	lock  sync.Mutex
}

func (l *Lazy[T]) doInit() {
	l.lock.Lock()
	if atomic.LoadInt32(&l.state) == 0 {
		l.v = l.init()
		l.init = nil
		atomic.StoreInt32(&l.state, 1)
	}
	l.lock.Unlock()
}

func (l *Lazy[T]) Get() T {
	if atomic.LoadInt32(&l.state) == 0 {
		l.doInit()
	}
	return l.v
}

func (l *Lazy[T]) GetRef() *T {
	if atomic.LoadInt32(&l.state) == 0 {
		l.doInit()
	}
	return &l.v
}

func NewLazy[T any](init func() T) Lazy[T] {
	return Lazy[T]{
		init: init,
	}
}
