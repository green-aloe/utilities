package pool

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Pool_Get tests that Pool's Get method returns the correct item from the pool, depending on
// the pool's state.
func Test_Pool_Get(t *testing.T) {
	t.Run("nil pool", func(t *testing.T) {
		var pool *Pool[int]
		item := pool.Get()
		require.Zero(t, item)
	})

	t.Run("zero pool", func(t *testing.T) {
		var pool Pool[uint16]
		item := pool.Get()
		require.Zero(t, item)
	})

	t.Run("empty pool, no callbacks", func(t *testing.T) {
		var pool Pool[int64]
		item := pool.Get()
		require.Zero(t, item)
	})

	t.Run("empty pool, NewItem callback", func(t *testing.T) {
		pool := Pool[float32]{
			NewItem: func() float32 { return 1.1 },
		}
		item := pool.Get()
		require.Equal(t, float32(1.1), item)
	})

	t.Run("empty pool, ClearItem callback", func(t *testing.T) {
		pool := Pool[float64]{
			ClearItem: func(float64) float64 { return 2.2 },
		}
		item := pool.Get()
		require.Zero(t, item)
	})

	t.Run("empty pool, both callbacks", func(t *testing.T) {
		pool := Pool[int8]{
			NewItem:   func() int8 { return 1 },
			ClearItem: func(int8) int8 { return 5 },
		}
		item := pool.Get()
		require.Equal(t, int8(1), item)
	})

	t.Run("non-empty pool, no callbacks", func(t *testing.T) {
		var pool Pool[rune]
		pool.Store('Z')

		item := pool.Get()
		require.Equal(t, 'Z', item)

		item = pool.Get()
		require.Zero(t, item)
	})

	t.Run("non-empty pool, NewItem callback", func(t *testing.T) {
		pool := Pool[string]{
			NewItem: func() string { return "test1" },
		}
		pool.Store("test2")

		item := pool.Get()
		require.Equal(t, "test2", item)

		item = pool.Get()
		require.Equal(t, "test1", item)
	})

	t.Run("non-empty pool, ClearItem callback", func(t *testing.T) {
		pool := Pool[bool]{
			ClearItem: func(bool) bool { return true },
		}
		pool.Store(true)

		item := pool.Get()
		require.Equal(t, true, item)

		item = pool.Get()
		require.Zero(t, item)
	})

	t.Run("non-empty pool, both callbacks", func(t *testing.T) {
		pool := Pool[uint64]{
			NewItem:   func() uint64 { return 1000 },
			ClearItem: func(uint64) uint64 { return 256 },
		}
		pool.Store(123)

		item := pool.Get()
		require.Equal(t, uint64(256), item)

		item = pool.Get()
		require.Equal(t, uint64(1000), item)
	})

	t.Run("concurrent use", func(t *testing.T) {
		pool := Pool[int]{
			NewItem:   func() int { return 1 },
			ClearItem: func(int) int { return 2 },
		}

		for i := 0; i < 100; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					item := pool.Get()
					require.Equal(t, 1, item)
				}
			}()
		}
	})
}

// Test_Pool_Store tests that Pool's Store method correctly store an object into the pool.
func Test_Pool_Store(t *testing.T) {
	t.Run("nil pool", func(t *testing.T) {
		var pool *Pool[float64]
		pool.Store(10.10)

		require.Equal(t, 0, pool.Count())
	})

	t.Run("zero pool", func(t *testing.T) {
		var pool Pool[rune]
		pool.Store('🫠')

		require.Equal(t, 1, pool.Count())
	})

	t.Run("empty pool, no callbacks", func(t *testing.T) {
		var pool Pool[string]
		pool.Store("test")

		require.Equal(t, 1, pool.Count())
		require.Equal(t, "test", pool.Get())
	})

	t.Run("empty pool, NewItem callback", func(t *testing.T) {
		pool := Pool[int32]{
			NewItem: func() int32 { return 2 },
		}
		pool.Store(4)

		require.Equal(t, 1, pool.Count())
		require.Equal(t, int32(4), pool.Get())
	})

	t.Run("empty pool, ClearItem callback", func(t *testing.T) {
		pool := Pool[int]{
			ClearItem: func(int) int { return 200 },
		}
		pool.Store(5)

		require.Equal(t, 1, pool.Count())
		require.Equal(t, 200, pool.Get())
	})

	t.Run("empty pool, both callbacks", func(t *testing.T) {
		pool := Pool[float32]{
			NewItem:   func() float32 { return 1.1 },
			ClearItem: func(float32) float32 { return 5.5 },
		}
		pool.Store(400)

		require.Equal(t, 1, pool.Count())
		require.Equal(t, float32(5.5), pool.Get())
	})

	t.Run("non-empty pool, no callbacks", func(t *testing.T) {
		var pool Pool[[]int]
		pool.Store([]int{1, 2, 3})
		pool.Store([]int{4, 5, 6})

		require.Equal(t, 2, pool.Count())
		require.Equal(t, []int{4, 5, 6}, pool.Get())
		require.Equal(t, []int{1, 2, 3}, pool.Get())
	})

	t.Run("non-empty pool, NewItem callback", func(t *testing.T) {
		pool := Pool[[]rune]{
			NewItem: func() []rune { return []rune{'a', 'b', 'c'} },
		}
		pool.Store([]rune{'d', 'e', 'f'})
		pool.Store([]rune{'g', 'h', 'i'})

		require.Equal(t, 2, pool.Count())
		require.Equal(t, []rune{'g', 'h', 'i'}, pool.Get())
		require.Equal(t, []rune{'d', 'e', 'f'}, pool.Get())
	})

	t.Run("non-empty pool, ClearItem callback", func(t *testing.T) {
		pool := Pool[[]rune]{
			ClearItem: func([]rune) []rune { return nil },
		}
		pool.Store([]rune{'d', 'e', 'f'})
		pool.Store([]rune{'g', 'h', 'i'})

		require.Equal(t, 2, pool.Count())
		require.Equal(t, []rune(nil), pool.Get())
		require.Equal(t, []rune(nil), pool.Get())
	})

	t.Run("non-empty pool, both callbacks", func(t *testing.T) {
		pool := Pool[int]{
			NewItem:   func() int { return 23 },
			ClearItem: func(int) int { return 24 },
		}
		pool.Store(123)
		pool.Store(321)

		require.Equal(t, 2, pool.Count())
		require.Equal(t, 24, pool.Get())
		require.Equal(t, 24, pool.Get())
	})

	t.Run("concurrent use", func(t *testing.T) {
		pool := Pool[int]{
			NewItem:   func() int { return 1 },
			ClearItem: func(int) int { return 2 },
		}

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				for j := 0; j < 100; j++ {
					pool.Store(j)
				}
			}()
		}
		wg.Wait()

		require.Equal(t, 10_000, pool.Count())
		for pool.Count() > 0 {
			require.Equal(t, 2, pool.Get())
		}
	})
}

// Test_Pool_Count tests that Pool's Count method returns the correct number of items in the pool.
func Test_Pool_Count(t *testing.T) {
	t.Run("nil pool", func(t *testing.T) {
		var pool *Pool[bool]
		require.Equal(t, 0, pool.Count())
	})

	t.Run("zero pool", func(t *testing.T) {
		var pool Pool[complex128]
		require.Equal(t, 0, pool.Count())
	})

	t.Run("empty pool", func(t *testing.T) {
		pool := Pool[int]{
			NewItem:   func() int { return 10 },
			ClearItem: func(int) int { return 20 },
		}
		require.Equal(t, 0, pool.Count())
	})

	t.Run("non-empty pool", func(t *testing.T) {
		pool := Pool[int]{
			NewItem:   func() int { return 10 },
			ClearItem: func(int) int { return 20 },
		}
		pool.Store(1)
		pool.Store(2)

		require.Equal(t, 2, pool.Count())

		pool.Get()
		require.Equal(t, 1, pool.Count())
		pool.Get()
		require.Equal(t, 0, pool.Count())
		pool.Get()
		require.Equal(t, 0, pool.Count())
	})

	t.Run("concurrent use", func(t *testing.T) {
		pool := Pool[int]{
			NewItem:   func() int { return 1 },
			ClearItem: func(int) int { return 2 },
		}
		for i := 0; i < 100; i++ {
			pool.Store(i)
		}

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				require.Equal(t, 100, pool.Count())
			}()
		}
		wg.Wait()
	})
}
