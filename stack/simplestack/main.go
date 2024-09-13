package simplestack

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

type Stack[T any] struct {
	items []T
}

func getZero[T Ordered]() T {
	var result T
	return result
}

func (stack *Stack[T]) Push(item T) {
	// if item != getZero[T](){
	stack.items = append(stack.items, item)
	// }
}

func (stack *Stack[T]) Pop() T {
	length := len(stack.items)
	// if length > 0 {
	returnValue := stack.items[length-1]
	stack.items = stack.items[:(length - 1)]
	return returnValue
	// } else {
	// 	return getZero[T]()
	// }
}

func (stack *Stack[T]) Top() T {
	length := len(stack.items)
	// if length > 0 {
	return stack.items[length-1]
	// } else {
	// 	return getZero[T]()
	// }
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func main() {
	// create a stack of name
	nameStack := Stack[string]{}
	nameStack.Push("Zachary")
	nameStack.Push("Adolf")

	if !nameStack.IsEmpty() {
		topOfStack := nameStack.Top()
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nPopped from stack is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nPopped from stack is %s", poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nPopped from stack is %s", poppedFromStack)
	}

	// create stack of intergers
	instStack := Stack[int]{}
	instStack.Push(5)
	instStack.Push(10)
	instStack.Push(0)

	if !instStack.IsEmpty() {
		top := instStack.Top()
		fmt.Printf("\nTop of intStack is %d", top)
	}

	if !instStack.IsEmpty() {
		popFromStack := instStack.Pop()
		fmt.Printf("\nValue popped from stack is %d", popFromStack)
	}

	if !instStack.IsEmpty() {
		popFromStack := instStack.Pop()
		fmt.Printf("\nValue popped from stack is %d", popFromStack)
	}

	if !instStack.IsEmpty() {
		popFromStack := instStack.Pop()
		fmt.Printf("\nValue popped from stack is %d", popFromStack)
	}

}
