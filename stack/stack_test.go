package stack

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Stack_Push tests that Stack's Push method adds a value to the top of the stack for various
// stack configurations.
func Test_Stack_Push(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack[int32]
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack[float32]
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
		require.NotPanics(t, func() { s.Push(1) })
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack[string]
		require.NotPanics(t, func() { s.Push("1") })
		require.NotPanics(t, func() { s.Push("2") })
		require.NotPanics(t, func() { s.Push("3") })
		require.Equal(t, 3, s.Count())
		require.Equal(t, "3", s.Pop())
		require.Equal(t, "2", s.Pop())
		require.Equal(t, "1", s.Pop())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack[int]

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				require.NotPanics(t, func() { s.Push(i) })
			}(i)
		}

		wg.Wait()
		require.Equal(t, 100, s.Count())
	})
}

// Test_Stack_Pop tests that Stack's Pop method removes and returns the value at the top of the
// stack for various stack configurations.
func Test_Stack_Pop(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack[int32]
		require.Zero(t, s.Pop())
		require.Zero(t, s.Pop())
		require.Zero(t, s.Pop())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack[rune]
		require.Zero(t, s.Pop())
		require.Zero(t, s.Pop())
		require.Zero(t, s.Pop())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack[int]
		for i := 1; i <= 10; i++ {
			s.Push(i)
		}
		for i := 10; i >= 1; i-- {
			require.Equal(t, i, s.Pop())
		}

		require.Zero(t, s.Pop())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack[int]

		var want []int
		for i := 0; i < 100; i++ {
			s.Push(i)
			want = append(want, i)
		}

		ch := make(chan int, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Pop()
			}()
		}

		var have []int
		for i := 0; i < 100; i++ {
			have = append(have, <-ch)
		}
		sort.Slice(have, func(i, j int) bool { return have[i] < have[j] })
		require.Equal(t, want, have)
		require.Len(t, ch, 0)

		require.Zero(t, s.Pop())
	})
}

// Test_Stack_CheckPop tests that Stack's CheckPop method returns the value at the top of the stack
// and a boolean indicating whether the stack is empty for various stack configurations.
func Test_Stack_CheckPop(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack[uint8]
		top, ok := s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)

		top, ok = s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)

		top, ok = s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack[string]
		top, ok := s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)

		top, ok = s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)

		top, ok = s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack[int]
		for i := 1; i <= 10; i++ {
			s.Push(i)
		}
		for i := 10; i >= 1; i-- {
			top, ok := s.CheckPop()
			require.Equal(t, i, top)
			require.True(t, ok)
		}

		top, ok := s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack[int]

		var want []int
		for i := 0; i < 100; i++ {
			s.Push(i)
			want = append(want, i)
		}

		ch := make(chan int, 100)
		for i := 0; i < 100; i++ {
			go func() {
				top, ok := s.CheckPop()
				require.True(t, ok)

				ch <- top
			}()
		}

		var have []int
		for i := 0; i < 100; i++ {
			have = append(have, <-ch)
		}
		sort.Slice(have, func(i, j int) bool { return have[i] < have[j] })
		require.Equal(t, want, have)
		require.Len(t, ch, 0)

		top, ok := s.CheckPop()
		require.Zero(t, top)
		require.False(t, ok)
	})
}

// Test_Stack_Peek tests that Stack's Peek method returns the value at the top of the stack without
// removing it for various stack configurations.
func Test_Stack_Peek(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack[[]rune]
		require.Zero(t, s.Peek())
		require.Zero(t, s.Peek())
		require.Zero(t, s.Peek())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack[[]any]
		require.Zero(t, s.Peek())
		require.Zero(t, s.Peek())
		require.Zero(t, s.Peek())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack[rune]
		s.Push('a')
		s.Push('b')
		s.Push('c')
		require.Equal(t, 'c', s.Peek())
		require.Equal(t, 'c', s.Peek())
		require.Equal(t, 'c', s.Peek())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack[float64]
		s.Push(1.1)
		s.Push(2.2)
		s.Push(3.3)

		ch := make(chan any, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Peek()
			}()
		}

		for i := 0; i < 100; i++ {
			v := <-ch
			f, ok := v.(float64)
			require.True(t, ok)
			require.Equal(t, 3.3, f)
		}
		require.Len(t, ch, 0)
	})
}

// Test_Stack_Empty tests that Stack's Empty method returns accurately determines if the stack is
// empty for various stack configurations.
func Test_Stack_Empty(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack[bool]
		require.True(t, s.Empty())
		require.True(t, s.Empty())
		require.True(t, s.Empty())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack[[]bool]
		require.True(t, s.Empty())
		require.True(t, s.Empty())
		require.True(t, s.Empty())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack[rune]
		for _, r := range "the quick brown fox jumps over the lazy dog" {
			s.Push(r)
			require.False(t, s.Empty())
		}
		require.False(t, s.Empty())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack[rune]

		ch := make(chan bool, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Empty()
			}()
		}

		for i := 0; i < 100; i++ {
			b := <-ch
			require.True(t, b)
		}
		require.Len(t, ch, 0)

		s.Push('😀')
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Empty()
			}()
		}

		for i := 0; i < 100; i++ {
			b := <-ch
			require.False(t, b)
		}
		require.Len(t, ch, 0)
	})
}

// Test_Stack_Count tests that Stack's Count method returns the correct number of elements in the
// stack for various stack configurations.
func Test_Stack_Count(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack[map[string]struct{}]
		require.Zero(t, s.Count())
		require.Zero(t, s.Count())
		require.Zero(t, s.Count())
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack[chan<- int]
		require.Zero(t, s.Count())
		require.Zero(t, s.Count())
		require.Zero(t, s.Count())
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack[string]
		s.Push("a")
		require.Equal(t, 1, s.Count())
		s.Push("b")
		require.Equal(t, 2, s.Count())
		s.Push("c")
		require.Equal(t, 3, s.Count())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack[int]

		ch := make(chan int, 100)
		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Count()
			}()
		}

		for i := 0; i < 100; i++ {
			i := <-ch
			require.Zero(t, i)
		}
		require.Len(t, ch, 0)

		for i := 0; i < 100; i++ {
			s.Push(i)
		}

		for i := 0; i < 100; i++ {
			go func() {
				ch <- s.Count()
			}()
		}

		for i := 0; i < 100; i++ {
			i := <-ch
			require.Equal(t, 100, i)
		}
		require.Len(t, ch, 0)
	})
}

// Test_Stack_Clear tests that Stack's Clear method removes all elements from the stack for various
// stack configurations.
func Test_Stack_Clear(t *testing.T) {
	t.Run("nil stack", func(t *testing.T) {
		var s *Stack[int32]
		require.NotPanics(t, func() { s.Clear() })
		require.True(t, s.Empty())

		require.NotPanics(t, func() { s.Clear() })
		require.NotPanics(t, func() { s.Clear() })
	})

	t.Run("empty stack", func(t *testing.T) {
		var s Stack[bool]
		require.NotPanics(t, func() { s.Clear() })
		require.True(t, s.Empty())

		require.NotPanics(t, func() { s.Clear() })
		require.NotPanics(t, func() { s.Clear() })
	})

	t.Run("non-empty stack", func(t *testing.T) {
		var s Stack[string]
		s.Push("a")
		s.Push("b")
		s.Push("c")
		require.False(t, s.Empty())
		require.NotPanics(t, func() { s.Clear() })
		require.True(t, s.Empty())
	})

	t.Run("concurrent use", func(t *testing.T) {
		var s Stack[int]

		for i := 0; i < 100; i++ {
			go func() {
				s.Clear()
			}()
		}

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.Push(1)
				s.Clear()
				s.Push(2)
				s.Clear()
				s.Push(3)
				s.Clear()
			}()
		}
		wg.Wait()
		require.True(t, s.Empty())
	})
}

// Test_differentTypes tests that Stack can handle values of different types.
func Test_differentTypes(t *testing.T) {
	var s Stack[any]
	s.Push(1)
	s.Push("2")
	s.Push(3.0)
	require.Equal(t, 3.0, s.Pop())
	require.Equal(t, "2", s.Pop())
	require.Equal(t, 1, s.Pop())
}

// Test_ConcurrentUse tests that various Stack methods can be used concurrently without panicking.
func Test_ConcurrentUse(t *testing.T) {
	for _, withClear := range []bool{false, true} {
		var s Stack[int]

		var wg sync.WaitGroup
		wg.Add(1_000_000)
		for i := 0; i < 1_000_000; i++ {
			go func(i int) {
				defer wg.Done()

				switch i % 16 {
				case 0, 1, 2, 3:
					s.Push(i)
				case 4, 5:
					s.Pop()
				case 6, 7:
					s.CheckPop()
				case 8, 9, 10:
					s.Peek()
				case 11, 12:
					s.Empty()
				case 13, 14:
					s.Count()
				case 15:
					if withClear {
						s.Clear()
					}
				}
			}(i)
		}
		wg.Wait()
	}
}

const (
	benchOpPush benchOp = iota
	benchOpPop
	benchOpCheckPop
	benchOpPeek
	benchOpEmpty
	benchOpCount
	benchOpClear
)

type benchOp int

func Benchmark_Push(b *testing.B) {
	benchmark_operation(b, benchOpPush, 1_000_000)
}
func Benchmark_Pop(b *testing.B) {
	benchmark_operation(b, benchOpPop, 1_000_000)
}
func Benchmark_CheckPop(b *testing.B) {
	benchmark_operation(b, benchOpCheckPop, 1_000_000)
}
func Benchmark_Peek(b *testing.B) {
	benchmark_operation(b, benchOpPeek, 1_000_000)
}
func Benchmark_Empty(b *testing.B) {
	benchmark_operation(b, benchOpEmpty, 1_000_000)
}
func Benchmark_Count(b *testing.B) {
	benchmark_operation(b, benchOpCount, 1_000_000)
}
func Benchmark_Clear(b *testing.B) {
	benchmark_operation(b, benchOpClear, 1_000_000)
}

func benchmark_operation(b *testing.B, op benchOp, size int) {
	var s Stack[int]
	for i := 1; i <= size; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch op {
		case benchOpPush:
			s.Push(i)
		case benchOpPop:
			s.Pop()
		case benchOpCheckPop:
			s.CheckPop()
		case benchOpPeek:
			s.Peek()
		case benchOpEmpty:
			s.Empty()
		case benchOpCount:
			s.Count()
		case benchOpClear:
			s.Clear()
		}
	}
}
