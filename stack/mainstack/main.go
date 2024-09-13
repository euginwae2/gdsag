package main

import (
	"fmt"
	"time"

	"example.com/nodestack"
	"example.com/simplestack"
)

const size = 10_000_000

func main() {
	nodeStack := nodestack.Stack[int]{}
	sliceStack := simplestack.Stack[int]{}

	// benchmark node stack

	start := time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Push(i)
	}
	elapsed := time.Since(start)
	fmt.Println("\nTime efor 10 million Push() operation nodeStack: ", elapsed)

	start = time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\n Time for 10 million Pop() operations on nodeStack: ", elapsed)

	// benchmark simpleStack
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Push(i)
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime efor 10 million Push() operation sliceStack: ", elapsed)

	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\n Time for 10 million Pop() operations on sliceStack: ", elapsed)

}
