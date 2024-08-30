package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

const LargestPrime = 10_000_000

var cores int
var primeNumbers []int
var m sync.Mutex
var wg sync.WaitGroup

func SieveOfEratosthenes(n int) []int {
	// Finds all
	primes := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		primes[i] = true
	}

	// The sieve logic
	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	// return all primes
	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}
	return primeNumbers
}

func primesBetween(prime []int, low, high int) {
	// computes the prime numbers betweeb low and high
	// given the initial set of primes from the SieveOferathosthenes
	defer wg.Done()
	limit := high - low
	segment := make([]bool, limit+1)
	for i := 0; i < len(segment); i++ {
		segment[i] = true
	}
	// find the primes on the current segment based on initial primes
	for i := 0; i < len(prime); i++ {
		lowlimit := int(math.Floor(float64(low)/float64(prime[i])) * float64(prime[i]))
		if lowlimit < low {
			lowlimit += prime[i]
		}

		for j := lowlimit; j < high; j += prime[i] {
			segment[j-low] = false
		}

		// each number in [low to high] is mapped to [0, high - low]
		for j := lowlimit; j < high; j += prime[i] {
			segment[j-low] = false
		}
	}

	m.Lock()
	for i := low; i < high; i++ {
		if segment[i-low] {
			primeNumbers = append(primeNumbers, i)
		}
	}
	m.Unlock()
}

func SegmentedSieve(n int) {
	limit := int(math.Floor(float64(n) / float64(cores)))
	prime := SieveOfEratosthenes(limit)
	for i := 0; i < len(prime); i++ {
		primeNumbers = append(primeNumbers, prime[i])
	}
	for low := limit; low < n; low += limit {
		high := low + limit
		if high >= n {
			high = n
		}
		wg.Add(1)
		go primesBetween(prime, low, high)
	}
	wg.Wait()
}

func main() {
	cores = runtime.NumCPU()
	start := time.Now()
	SegmentedSieve(LargestPrime)
	elapsed := time.Since(start)
	fmt.Println("\nNumber of cores: ", cores)
	fmt.Println("\nComputation time for concurrent: ", elapsed)
	fmt.Println("Number of primes = ", len(primeNumbers))
}
