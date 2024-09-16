package main

import (
	"fmt"

	"example.com/simplestack"
)

func convertToBinary(input int) (binary []int) {
	binaryNumberStack := simplestack.Stack[int]{}
	for {
		binaryNumberStack.Push(input % 2)
		input = input / 2
		if input == 0 {
			break
		}
	}
	binary = []int{}
	for {
		if !binaryNumberStack.IsEmpty() {
			binary = append(binary, binaryNumberStack.Pop())
		} else {
			break
		}
	}
	return binary
}

func main() {
	number := 1_000_000
	binaryNumber := convertToBinary(number)
	fmt.Printf("\n%d converted to binary is \n %v", number, binaryNumber)
}
