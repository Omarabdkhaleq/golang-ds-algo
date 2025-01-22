package stack

import "fmt"

/* STACK => LIFO */

type Stack struct {
	top   int
	stack []int
}

func NewStack(size int) Stack {
	return Stack{
		top:   -1,
		stack: make([]int, 0, size),
	}
}

func (s Stack) Push(i int) Stack {
	if s.IsFull() {
		panic("stack is full")
	}
	s.stack = append(s.stack, i)
	s.top++
	return s
}

func (s Stack) Pop() Stack {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	s.stack = s.stack[:s.top]
	s.top--
	return s
}

func (s Stack) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func (s Stack) IsFull() bool {
	if len(s.stack) == 0 {
		return false
	}
	if s.top == len(s.stack) {
		return true
	}
	return false
}

func (s Stack) Peek() int {
	return s.stack[s.top]
}

func FromScratch() {
	s := NewStack(4)
	s = s.Push(1)
	fmt.Println(&s.stack)
	s = s.Push(2)
	fmt.Println(&s.stack)
	fmt.Println(s.Peek())
	s = s.Pop()
	fmt.Println(&s.stack)
	s = s.Pop()
	fmt.Println(&s)
}
