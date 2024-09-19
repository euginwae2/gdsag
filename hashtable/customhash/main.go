package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"time"
)

const tableSize = 100_000
var length int

func hash(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

type WordType struct {
	word string
	list []string
}

type HashTable [tableSize]WordType

func Newtable() HashTable {
	var table HashTable
	for i := 0; i < tableSize; i++ {
		table[i] = WordType{"", []string{}}
	}
	return table
}

func (table *HashTable) Insert(word string) {
	index := hash(word) % tableSize //Between 0 and tableSize -1
	// Search table[index] for word
	if table[index].word == word {
		return // duplicates not allowed
	}
	if len(table[index].list) > 0 {
		for i := 0; i < len(table[index].list); i++ {
			if table[index].list[i] == word {
				return //duplicates not allowed
			}
		}
	}
	if table[index].word == "" {
		table[index].word = word
	} else {
		table[index].list = append(table[index].list, word)
	}
	length += 1
}

func (table *HashTable) IsPresent(word string) bool {
	index := hash(word) % tableSize
	if table[index].word == word {
		return true
	}

	if len(table[index].list) > 0 {
		for i := 0; i < len(table[index].list); i++ {
			if table[index].list[i] == word {
				return true
			}
		}
	}
	return false

}


func main() {
	myTable := Newtable()
	mapCollection := make(map[string]string)

	words := []string{}

	for i :=0; i <500_000; i++ {
		word := strconv.Itoa(i)
		words = append(words, word)
		myTable.Insert(word)
		mapCollection[word] = ""
	}

	fmt.Println("Benchmark test begins to test words: ", length)
	start := time.Now()
	for i :=0; i <length; i++ {
		if !myTable.IsPresent(words[i]) {
			fmt.Println("Word not found in table: ", words[i])
		}
	}
	elapsed := time.Since(start)

	fmt.Println("Time to test all words in myTable: ", elapsed)

	start = time.Now()
	for i :=0; i<length; i++ {
		_,present := mapCollection[words[i]]
		if !present {
			fmt.Println("Word not found in mapCollection: ", words[i])
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Time to test all words in mapCollection: ", elapsed)


}
