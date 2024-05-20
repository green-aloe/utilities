package stack

import "sync"

// A stack is a first-in-last-out (FILO) data structure: the last element pushed is the first
// element popped. The zero value is an empty stack and ready to use, and stacks are safe for
// concurrent use.
type Stack[T any] struct {
	top   *node[T]
	mutex sync.Mutex
}

// A node is a single element in a stack.
type node[T any] struct {
	v    T
	next *node[T]
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(v T) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.top = &node[T]{
		v:    v,
		next: s.top,
	}
}

// Pop removes and returns the value at the top of the stack. If the stack is empty, this returns
// the zero value of the stack's type.
func (s *Stack[T]) Pop() (t T) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.top == nil {
		return
	}

	v := s.top.v
	s.top = s.top.next

	return v
}

// CheckPop returns the value at the top of the stack and a boolean indicating whether the stack is
// empty. If the stack is empty, this returns the zero value of the stack's type and false.
func (s *Stack[T]) CheckPop() (t T, ok bool) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.top == nil {
		return
	}

	v := s.top.v
	s.top = s.top.next

	return v, true
}

// Peek returns the value at the top of the stack without removing it. If the stack is empty, this
// returns the zero value of the stack's type.
func (s *Stack[T]) Peek() (t T) {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.top == nil {
		return
	}

	return s.top.v
}

// Empty returns true if the stack is empty.
func (s *Stack[T]) Empty() bool {
	if s == nil {
		return true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.top == nil
}

// Count returns the number of elements in the stack.
func (s *Stack[T]) Count() int {
	if s == nil {
		return 0
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	count := 0
	for n := s.top; n != nil; n = n.next {
		count++
	}

	return count
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	if s == nil {
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.top = nil
}
