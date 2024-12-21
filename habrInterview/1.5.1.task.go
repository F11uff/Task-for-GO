package main

import "fmt"

//Реализовать структуру данных "стек" с функциональностью pop, append и top.

type Stack struct {
	items []int
}

func main() {
	stack := Stack{}
	stack.Push(1)
	fmt.Println(stack.items)
	stack.Push(2)
	fmt.Println(stack.items)
	stack.Push(3)
	fmt.Println(stack.items)
	stack.Pop()
	fmt.Println(stack.items)
	stack.Pop()
	fmt.Println(stack.items)
	stack.Pop()
	fmt.Println(stack.items)
	stack.Pop()
	fmt.Println(stack.items)
	stack.Pop()
	fmt.Println(stack.items)
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		return
	}

	s.items = s.items[:len(s.items)-1]
}
