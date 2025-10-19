package queue

import "sync"

type Queue[T any] struct {
	items chan T
	once  sync.Once
}

func New[T any](size int) *Queue[T] {
	return &Queue[T]{items: make(chan T, size)}
}

func (q *Queue[T]) Push(v T) { q.items <- v }
func (q *Queue[T]) Pop() (T, bool) {
	v, ok := <-q.items
	return v, ok
}
func (q *Queue[T]) Chan() <-chan T { return q.items }
func (q *Queue[T]) Close() {
	q.once.Do(func() { close(q.items) })
}
