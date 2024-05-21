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
