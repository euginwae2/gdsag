package nodestack

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	first *Node[T]
}

// Methods
// push
func (s *Stack[T]) Push(item T) {
	newNode := Node[T]{item, nil}
	newNode.next = s.first
	s.first = &newNode
}

// top
func (s *Stack[T]) Top() T {
	return s.first.value
}

// pop
func (s *Stack[T]) Pop() T {
	result := s.first.value
	s.first = s.first.next
	return result
}

// isEmpty
func (s *Stack[T]) IsEmpty() bool {
	return s.first == nil
}

func main() {
	nameStack := Stack[string]{}
	nameStack.Push("Zachary")
	nameStack.Push("Adolf")

	if !nameStack.IsEmpty() {
		topOfStack := nameStack.Top()
		fmt.Printf("\n TopOfStack value is %s", topOfStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\n Popped form Stack value is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\n Popped form Stack value is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\n Popped form Stack value is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\n Popped form Stack value is %s", poppedFromStack)
	}

	intStack := Stack[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0)

	if !intStack.IsEmpty() {
		topOfStack := intStack.Top()
		fmt.Printf("\n Top of intStack is %d", topOfStack)
	}

	if !intStack.IsEmpty() {
		poppedFromStack := intStack.Pop()
		fmt.Printf("\nPoppedFromStack value is %d", poppedFromStack)
	}

	if !intStack.IsEmpty() {
		poppedFromStack := intStack.Pop()
		fmt.Printf("\nPoppedFromStack value is %d", poppedFromStack)
	}

	if !intStack.IsEmpty() {
		poppedFromStack := intStack.Pop()
		fmt.Printf("\nPoppedFromStack value is %d", poppedFromStack)
	}

	if !intStack.IsEmpty() {
		poppedFromStack := intStack.Pop()
		fmt.Printf("\nPoppedFromStack value is %d", poppedFromStack)
	}

}
