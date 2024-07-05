package pool

import (
	"github.com/green-aloe/utility/stack"
)

// A pool is a collection of items that can be reused. If a pool is empty when an item is requested,
// it can generate a new item. The zero value of a pool is ready to use and safe for concurrent
// access by multiple goroutines.
//
// One of the key differences between this pool and a sync.Pool is that this pool does not
// automatically remove any items stored in it and has no automatic cleanup mechanism.
type Pool[T any] struct {
	// NewItem generates a new item when the pool is empty.
	NewItem func() T
	// PreStore is called before storing an item in the pool. It enables additional functionality
	// like monitoring or transforming items as they are stored.
	PreStore func(T) T

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

// Store stores an object in the pool for later reuse. If PreStore is non-nil, the pool clears the
// item before storing it.
func (pool *Pool[T]) Store(t T) {
	if pool == nil {
		return
	}

	if pool.PreStore != nil {
		t = pool.PreStore(t)
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

// Clear removes all items from the pool.
func (pool *Pool[T]) Clear() {
	if pool == nil {
		return
	}

	pool.stack.Clear()
}
