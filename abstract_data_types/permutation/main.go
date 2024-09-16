package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func init() {
	buildDictionary()
}

var dictionary map[string][]string

func alphabetize(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func buildDictionary() {
	dictionary = make(map[string][]string)

	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textwords []string

	for scanner.Scan() {
		textwords = append(textwords, scanner.Text())
	}

	file.Close()

	for _, word := range textwords {
		alphabetized := alphabetize(word)
		var lst []string

		if len(dictionary) > 0 && len(dictionary[alphabetized]) > 0 {
			lst = dictionary[alphabetized]
		} else {
			lst = []string{}
		}
		lst = append(lst, word)
		dictionary[alphabetized] = lst
	}
}

func output(word string) {
	wd := alphabetize(word)
	fmt.Printf("Permutatuin group for %s is %s", word, dictionary[wd])
}

func main() {
	output("parties")
}
