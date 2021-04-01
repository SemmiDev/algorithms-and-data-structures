package main

import "fmt"

type Stack []interface{}

type StackOperations interface {
	IsEmpty() bool
	size() int
	Push(value interface{})
	Delete()
	Pop() interface{}
	Peek() interface{}
}

func NewStack() StackOperations {
	return &Stack{}
}

func (s *Stack) size() int {
	return len(*s)
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Delete() {
	if s.IsEmpty() {
		return
	}

	index := len(*s) - 1
	*s = (*s)[:index]
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func (s *Stack) Peek() interface{} {
	latestIndex := len(*s) - 1
	return (*s)[latestIndex]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	stack := NewStack()

	fmt.Println(stack.IsEmpty() == true)

	stack.Push("one")
	stack.Push(2)
	stack.Push("three")
	stack.Push(4)
	stack.Push(true)
	stack.Push(false)
	stack.Push(7)
	stack.Push(8)
	stack.Push(9)
	stack.Push("ten")

	fmt.Println(stack.IsEmpty() == false)
	fmt.Println(stack.Peek() == "ten")
	fmt.Println(stack.size() == 10)
	stack.Delete()
	fmt.Println(stack.size() == 9)
	fmt.Println(stack.Peek() == 9)
	fmt.Println(stack.Pop() == 9)
	fmt.Println(stack.size() == 8)
}