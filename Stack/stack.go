package main

import "fmt"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() int {
	l := len(*s)
	v := (*s)[l-1]
	return v
}

func Main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
}
