package main

import (
	"fmt"

	"example.com/counter"
)

func main() {
	myCounter := counter.Counter{}

	myCounter.Increment()
	myCounter.Increment()
	myCounter.Reset()
	myCounter.Increment()
	myCounter.Increment()
	myCounter.Increment()
	myCounter.Increment()
	myCounter.Decrement()
	countValue := myCounter.GetCount()
	fmt.Println(countValue)
}
