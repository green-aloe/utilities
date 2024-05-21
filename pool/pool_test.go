package pool

import (
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
