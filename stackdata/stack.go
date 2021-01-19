package main

import "fmt"

// Stack reperesnts stack that hold the slice
type Stack struct {
	items []int
}

// Push will add the value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and RETURNS the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(10)
	myStack.Push(10)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)

}
