package main

import (
	"fmt"
	"math"
	"time"
)

const LargestPrime = 10_000_000

var cores int

func SieveOfEratosthenes(n int) []int {
	// find all primes up to n
	primes := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		primes[i] = true
	}

	// the sieve logic
	for p := 2; p*p <= n; p++ {
		if primes[p] {
			// update a;; miltiples of p
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}

	return primeNumbers
}

func primesBetween(prime []int, low, high int) []int {
	// Computes the prime number between low and high
	// given the initial set of primes from the SieveOdEratosthenes
	limit := high - low
	var result []int
	segment := make([]bool, limit+1)
	for i := 0; i < len(segment); i++ {
		segment[i] = true
	}

	// Find the primes in the current segment bnased on initial primes
	for i := 0; i < len(prime); i++ {
		lowlimit := int(math.Floor(float64(low)/float64(prime[i])) * float64(prime[i]))

		if lowlimit < low {
			lowlimit += prime[i]
		}

		for j := lowlimit; j < high; j += prime[i] {
			segment[j-low] = false
		}
	}

	for i := low; i < high; i++ {
		if segment[i-low] {
			result = append(result, i)
		}
	}

	return result
}

func SegementedSieve(n int) []int {
	// each segment is of size square root of n
	// find all primes up to n
	var primeNumbers []int
	limit := (int)(math.Floor(math.Sqrt((float64(n)))))

	prime := SieveOfEratosthenes(limit)
	for i := 0; i < len(prime); i++ {
		primeNumbers = append(primeNumbers, prime[i])
	}

	low := limit
	high := 2 * limit
	if high >= n {
		high = n
	}

	for {
		if low < n {
			next := primesBetween(prime, low, high)
			// fmt.Printf("\nprimesBetween(%d, %d) = %v", low, high, next)
			for i := 0; i < len(next); i++ {
				primeNumbers = append(primeNumbers, next[i])
			}
			low = low + limit
			high = high + limit
			if high >= n {
				high = n
			}
		} else {
			break
		}
	}
	return primeNumbers
}

func main() {
	start := time.Now()
	primeNumbers := SegementedSieve(LargestPrime)
	elapsed := time.Since(start)
	fmt.Println("\nComputation time: ", elapsed)
	fmt.Println("Number of primes = ", len(primeNumbers))
}
