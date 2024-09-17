package main

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/deque"
)


func MaxSubarrayBruteForce(items []int, k int) []int{
	res := []int{}
	for i := 0; i <= len(items)-k; i++ {
		max := items[i]
		for j := 0; j < k; j++ {
			if items[j +i] > max {
				max = items[j + i]
			}
		}
		res = append(res,max)
	}
	return res
}

func MaxSubArrayUsingDeque(input []int, k int) (output []int) {

	deque := deque.Deque[int]{}
	var index int
	// first window
	for index = 0; index < k; index++ {
		for {
			if deque.Empty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}

	for; index < len(input); index ++ {
		output =  append(output, input[deque.First()])

		// Remove elements out of the window
		for {
			if deque.Empty() || deque.First() > index -k {
				break
			}
			deque.RemoveFirst()
		}
		// Remove values smaller than the element currently being added
		for {
			if deque.Empty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}
	output = append(output, input[deque.First()])
	return output
}

const size = 1_000_000
func main() {
	input := []int{9,1,1,0,0,0,1,0,6,8}
	// input := []int{3,1,6,4,2,10,5,9}
	output1 := MaxSubarrayBruteForce(input,3)
	fmt.Println("output = ", output1)
	output2 := MaxSubArrayUsingDeque(input,3)
	fmt.Println("output = ", output2)

	// Benchmark performance of two algorithms
	input = []int{}
	for i := 0; i < size; i++ {
		input =  append(input, rand.Intn(1000))
	}
	start := time.Now()
	MaxSubArrayUsingDeque(input,10000)
	elapsed := time.Since(start)
	fmt.Println("Using Deque: ", elapsed)

	start =  time.Now()
	MaxSubarrayBruteForce(input, 10000)
	elapsed =  time.Since(start)
	fmt.Println("using Brute Force: ", elapsed)
}