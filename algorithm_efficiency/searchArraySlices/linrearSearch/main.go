package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const size = 100_000_000

type Ordered interface {
	~int | ~float64 | ~string
}

func searchSegment[T Ordered](slice []T, target T, a, b int, ch chan<- bool) {
	// Generates boolean value put into ch
	for i := a; i < b; i++ {
		if slice[i] == target {
			ch <- true
		}
	}
	ch <- false
}

func concurrentSearch[T Ordered](data []T, target T) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(len(data)) / float64(numSegments))
	// launch numsegments goroutines
	for index := 0; index < numSegments; index++ {
		go searchSegment(data, target, index*segmentSize, index*segmentSize+segmentSize, ch)
	}
	num := 0 //Completed goroutines
	for {
		select {
		case value := <-ch: //Blocks until a goroutine puts a bool into the channel
			if value == true {
				return true
			}
			num += 1
			if num == numSegments { //All goroutines have completed
				return false
			}
		}
	}
	return false
}

func linearSearch[T Ordered](slice []T, target T) bool {
	// return true if target is in the slice
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false
}

func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	start := time.Now()
	result := concurrentSearch[float64](data, 54.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_00_00 floats using linear search  = ", elapsed)
	fmt.Println("Results of search is ", result)

	start = time.Now()
	result = concurrentSearch[float64](data, data[size/2])
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using linear search = ", elapsed)
	fmt.Println("Result of search is ", result)
}
