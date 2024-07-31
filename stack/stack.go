package stack

import "sync"

// A stack is a first-in-last-out (FILO) data structure: the last element pushed onto the stack is
// the first one popped from it. The zero value is an empty stack and ready to use. A stack is safe
// for concurrent use.
type Stack[T any] struct {
	items []T
	mutex sync.Mutex
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(v T) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = append(s.items, v)
}

// Pop removes and returns the value at the top of the stack. If the stack is empty, this returns
// the zero value of the stack's type.
func (s *Stack[T]) Pop() (t T) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		return
	}

	t, s.items[len(s.items)-1] = s.items[len(s.items)-1], t
	s.items = s.items[:len(s.items)-1]

	return t
}

// CheckPop returns the value at the top of the stack and a boolean indicating whether the stack is
// empty. If the stack is empty, this returns the zero value of the stack's type and false.
func (s *Stack[T]) CheckPop() (t T, ok bool) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		return
	}

	t, s.items[len(s.items)-1] = s.items[len(s.items)-1], t
	s.items = s.items[:len(s.items)-1]

	return t, true
}

// Peek returns the value at the top of the stack without removing it. If the stack is empty, this
// returns the zero value of the stack's type.
func (s *Stack[T]) Peek() (t T) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		return
	}

	return s.items[len(s.items)-1]
}

// Empty returns true if the stack is empty.
func (s *Stack[T]) Empty() bool {
	if s == nil {
		return true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.items) == 0
}

// Count returns the number of elements in the stack.
func (s *Stack[T]) Count() int {
	if s == nil {
		return 0
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.items)
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = nil
}
