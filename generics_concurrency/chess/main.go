package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quit chan bool

func main() {
	// rand.Seed(time.Now().UnixNano())
	move := make(chan int)
	quit = make(chan bool)

	// launch two players
	go player("Bobby Fischer", move)
	go player("Boris Spassky", move)

	// start move
	move <- 1
	<-quit //Blocks until quit assigned a value
}

// player simulates a person maving in chess
func player(name string, move chan int) {
	// This function takes data out of the move channel
	// and puts data back into the move channel
	for {
		// wait for turn to play
		turn := <-move
		n := rand.Intn(100)
		if n <= 5 && turn >= 5 {
			fmt.Printf("Player %s was check mated and loses!", name)
			quit <- true
			return
		}

		// Display and then increment the total move count by one
		fmt.Printf("Player %s has moved. Turn %d\n", name, turn)
		turn++

		// Yield the turn back to the opposing player
		time.Sleep(1 * time.Second)
		move <- turn
	}
}
