package main

import (
	"fmt"
	"time"
)

const (
	NumbersToSum = 10_000_000
)

func sum(s []float64, c chan<- float64) {
	// A generator that puts data into channel
	sum := 0.0
	for _, v := range s {
		sum += float64(v)
	}
	c <- sum //blocks until c is taken out of the channel
}

func plainSum(s []float64) float64 {
	sum := 0.0
	for _,v := range s{
		sum += float64(v)
	}
	return sum
}

func main() {
	s := []float64{}
	for i := 0; i < NumbersToSum; i++ {
		s = append(s, 1.0)
	}

	c := make(chan float64)
	start := time.Now()
	go sum(s[:len(s) /2], c)
	go sum(s[len(s) /2 :], c)
	first,second := <-c, <-c //receive from each c
	elapsed := time.Since(start)
	fmt.Printf("first: %f second: %f elapsed time: %v", first,second, elapsed)

	start =  time.Now()
	answer := plainSum(s)
	elapsed = time.Since(start)
	fmt.Printf("\nplain sum: %f elapsed time: %v\n", answer,elapsed)
}