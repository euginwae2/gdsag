package main

import (
	"fmt"
	"strings"
	"time"
)

// rabin karp
// reduce search from O(n*m) -> O(n)

func BruteForseSearch(txt, pattern string) (bool, int) {
	patternLenth := len(pattern)
	for outer := 0; outer < len(txt)-patternLenth; outer++ {
		if txt[(outer):(outer+patternLenth)] == pattern {
			return true, outer
		}
	}
	return false, -1
}

const (
	Radix = uint64(10)
	Q     = uint64(10 ^ 9 + 9)
)

func Hash(s string, Length int) uint64 {
	// Horner's method
	h := uint64(0)
	for i := 0; i < Length; i++ {
		h = (h*Radix + uint64(s[i])) % Q
	}
	return h
}

func Search(txt, pattern string) (bool, int) {
	strings.ToLower(txt)
	strings.ToLower(pattern)
	n := len(txt)
	m := len(pattern)

	patternHash := Hash(pattern, m)
	textHash := Hash(txt, m)
	if textHash == patternHash {
		return true, 0
	}
	PM := uint64(1)
	for i := 1; i <= m-1; i++ {
		PM = (Radix * PM) % Q
	}
	for i := m; i < n; i++ {
		textHash = (textHash + Q - PM*uint64(txt[i-m])%Q) % Q
		textHash = (textHash*Radix + uint64(txt[i])) % Q
		if (patternHash == textHash) && pattern == txt[(i-m+1):(i+1)] {
			return true, i - m + 1
		}
	}
	return false, -1
}

func main() {
	text := "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	pattern := "816406286208998628034825342"

	start := time.Now()
	_,_ = BruteForseSearch(text,pattern)
	elapsed :=time.Since(start)
	fmt.Println("Computation time using BruteforceSearch: ", elapsed)

	start = time.Now()
	_,_ = Search(text,pattern)
	elapsed =  time.Since(start)
	fmt.Println("Computation time using Rabin Karp: ", elapsed)

	fmt.Println(BruteForseSearch(text,pattern))
	fmt.Println(Search(text,pattern))
}
