package main

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

type Stack[T Ordered] struct {
	items []T
}

func getZero[T Ordered]() T{
	var result T
	return result
}

func (stack *Stack[T]) Push(item T) {
	if item != getZero[T](){
		stack.items = append(stack.items, item)
	}
}

func (stack *Stack[T]) Pop() T {
	length := len(stack.items)
	if length > 0 {
		returnValue := stack.items[length-1]
		stack.items =  stack.items[:(length -1)]
		return returnValue
	} else {
		return getZero[T]()
	}
}

func (stack *Stack[T]) Top() T{
	length := len(stack.items)
	if length > 0 {
		return stack.items[length-1]
	} else {
		return getZero[T]()
	}
}


func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func main(){
	// create a stack of name
	nameStack := Stack[string]{}
	nameStack.Push("Zachary")
	nameStack.Push("Adolf")
	topOfStack := nameStack.Top()
	if topOfStack != getZero[string]() {
		fmt.Printf("\nTop of stack is %s",topOfStack)
	}
	
	poppedFromStack := nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nPopped from stack is %s", poppedFromStack)
	}
	
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nPopped from stack is %s", poppedFromStack)
	}

	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nPopped from stack is %s", poppedFromStack)
	}


	// create stack of intergers
	 instStack := Stack[int]{}
	 instStack.Push(5)
	 instStack.Push(10)
	 instStack.Push(0)

	 top := instStack.Top()
	 if top != getZero[int]() {
		fmt.Printf("\nTop of intStack is %d", top)
	 }

	 popFromStack := instStack.Pop()
	 if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from stack is %d", popFromStack)
	 }

	 popFromStack = instStack.Pop()
	 if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from stack is %d", popFromStack)
	 }

	 popFromStack = instStack.Pop()
	 if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped from stack is %d", popFromStack)
	 }


}