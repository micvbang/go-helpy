package syncy

import "sync"

// Pool is a simple wrapper around sync.Pool, providing type safety.
type Pool[T any] struct {
	pool *sync.Pool
}

func NewPool[T any](new func() T) *Pool[T] {
	pool := &sync.Pool{
		New: func() any {
			return new()
		},
	}

	return &Pool[T]{
		pool: pool,
	}
}

func (p *Pool[T]) Put(v T) {
	p.pool.Put(v)
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}
