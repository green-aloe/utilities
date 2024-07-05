package stack_test

import (
	"fmt"
	"math"

	"github.com/green-aloe/utilities/stack"
)

func ExampleStack_Push() {
	var s stack.Stack[int]
	s.Push(123)

	count := s.Count()
	top := s.Peek()

	fmt.Println(count, top)

	// Output:
	// 1 123
}

func ExampleStack_Pop() {
	var s stack.Stack[string]
	top1 := s.Pop()

	s.Push("hello")
	s.Push("world")

	top2 := s.Pop()
	top3 := s.Pop()
	top4 := s.Pop()

	fmt.Println(top1)
	fmt.Println(top2)
	fmt.Println(top3)
	fmt.Println(top4)

	// Output:
	//
	// world
	// hello
	//
}

func ExampleStack_CheckPop() {
	var s stack.Stack[uint8]
	top1, ok1 := s.CheckPop()

	s.Push(3)
	s.Push(5)

	top2, ok2 := s.CheckPop()
	top3, ok3 := s.CheckPop()
	top4, ok4 := s.CheckPop()

	fmt.Println(top1, ok1)
	fmt.Println(top2, ok2)
	fmt.Println(top3, ok3)
	fmt.Println(top4, ok4)

	// Output:
	// 0 false
	// 5 true
	// 3 true
	// 0 false
}

func ExampleStack_Peek() {
	var s stack.Stack[[]string]
	top1 := s.Peek()

	s.Push([]string{"a", "b", "c"})
	s.Push([]string{"d", "e", "f"})

	top2 := s.Peek()
	top3 := s.Peek()

	fmt.Println(top1)
	fmt.Println(top2)
	fmt.Println(top3)

	// Output:
	// []
	// [d e f]
	// [d e f]
}

func ExampleStack_Empty() {
	var s stack.Stack[float64]
	isEmpty1 := s.Empty()

	s.Push(math.Pi)
	isEmpty2 := s.Empty()

	s.Pop()
	isEmpty3 := s.Empty()

	fmt.Println(isEmpty1, isEmpty2, isEmpty3)

	// Output:
	// true false true
}

func ExampleStack_Count() {
	var s stack.Stack[byte]
	fmt.Println(s.Count())

	for _, b := range []byte{'a', 'b', 'c'} {
		s.Push(b)
		fmt.Println(s.Count())
	}

	s.Pop()
	fmt.Println(s.Count())

	s.Peek()
	fmt.Println(s.Count())

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 2
	// 2
}

func ExampleStack_Clear() {
	var s stack.Stack[bool]
	s.Push(true)
	s.Push(false)
	count1 := s.Count()

	s.Clear()
	count2 := s.Count()

	fmt.Println(count1, count2)

	// Output:
	// 2 0
}
