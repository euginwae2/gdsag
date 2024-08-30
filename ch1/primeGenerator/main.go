package main

import (
	"fmt"
	"time"
)

const LargestPrime = 10_000_000

func SieveOferatosthenes(n int) []int {
	// finds all primes up to n
	primes := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		primes[i] = true
	}

	// The sieve logic for removing non-prime indices
	for p := 2; p*p <= n; p++ {
		if primes[p] {
			// update all multiples of p
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	// return all prim numbers <= n
	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}
	return primeNumbers
}


func main() {
	start := time.Now()
	sieve := SieveOferatosthenes(LargestPrime)
	elapsed := time.Since(start)
	fmt.Println("\nComputation time: ", elapsed)
	fmt.Println("Largest prime: ", sieve[len(sieve)-1])
	fmt.Println("\nNumber of primes: ", len(sieve))
}