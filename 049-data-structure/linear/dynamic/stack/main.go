package main

import "fmt"

// Stack and Queue
// Difference is how you take data out

// Stacks are LIFO (Last in first out)
// Stacks use PUSH and POP

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() {
	if len(s.items) > 0 {
		s.items = s.items[:len(s.items)-1]
	}
}

func main() {

	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)

}
