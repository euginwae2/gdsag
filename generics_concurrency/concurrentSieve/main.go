package main

import (
	"fmt"
	"time"
)

const LargestPrime = 100_000

var primes []int

// send the sequence to channel
func generate(prime1 chan<- int) {
	for i := 3; ; i += 2 {
		prime1 <- i //send 'i' to channel prime1
	}
}

// copy the values from channel 'in' to channel 'out'
// remove those divisible by primes
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in //Receive value from 'in'
		if i%prime != 0 {
			out <- i //send 'i' to 'out'
		}
	}
}

func main() {
	start := time.Now()
	prime1 := make(chan int) //create a new channel
	go generate(prime1)      //launch goroutine
	for {
		prime := <-prime1
		if prime > LargestPrime {
			break
		}
		primes = append(primes, prime)
		prime2 := make(chan int)
		go Filter(prime1, prime2, prime)
		prime1 = prime2
	}

	elapsed := time.Since(start)
	fmt.Println("computation time ", elapsed)
	fmt.Println("Number of primes = ", len(primes))
}
