package pool

import (
	"github.com/green-aloe/utility/stack"
)

type Pool[T any] struct {
	// NewItem generates a new item when the pool is empty.
	NewItem func() T
	// ClearItem clears an item before storing it in the pool.
	ClearItem func(T) T

	stack stack.Stack[T]
}

// Get returns a new or recycled object from the pool. If pool is nil or pool is empty and NewItem
// is nil, this returns the zero value of T.
func (pool *Pool[T]) Get() (t T) {
	if pool == nil {
		return
	}

	if t, ok := pool.stack.CheckPop(); ok {
		return t
	}

	if pool.NewItem != nil {
		return pool.NewItem()
	}

	return
}

// Store stores an object in the pool for later reuse. If ClearItem is non-nil, the pool clears the
// item before storing it.
func (pool *Pool[T]) Store(t T) {
	if pool == nil {
		return
	}

	if pool.ClearItem != nil {
		t = pool.ClearItem(t)
	}
	pool.stack.Push(t)
}

// Count returns the number of items in the pool.
func (pool *Pool[T]) Count() int {
	if pool == nil {
		return 0
	}

	return pool.stack.Count()
}
