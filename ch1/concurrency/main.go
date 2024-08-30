package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// simple goroutine running concurrent with main

func regularFunction(){
	fmt.Println("Just entered regularFunction()")
	time.Sleep(10 * time.Second)
}

func goroutineFunction(){
	fmt.Println("Just entered goroutineFunction()")
	time.Sleep(5 * time.Second)
	fmt.Println("goroutineFunction finished its work")
}

var wg sync.WaitGroup
func outputStrings() {
	defer wg.Done()
	strings := [5]string{"One", "Two", "Three", "Four", "Five"}
	for i := 0; i< 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(strings[i])
	}
}

func outputInts(){
	defer wg.Done()
	for i := 0; i < 5; i++ {
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(i)
	}
}

func outputFloats() {
	defer wg.Done()
	for i := 0; i<5;i++{
		delay := 1 + rand.Intn(3)
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println(float64(i *i) + 0.5)
	}
}
func main() {
	// go goroutineFunction()
	// fmt.Println("In main one line below goroutineFunction()")
	// regularFunction()
	// fmt.Println("In main one line below regularFunction()")

	wg.Add(3)
	go outputStrings()
	go outputInts()
	go outputFloats()
	wg.Wait()
}