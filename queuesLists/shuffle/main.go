package main

import (
	"fmt"

	"math/rand"

	"example.com/nodequeue"
)


type Card struct {
	Rank string
	Suit string
}

type Deck struct {
	Cards []Card
}

var ranks = []string{"2","3","4", "5", "6", "7", "8", "9", "10", "J","Q", "K", "A"}
var suits = []rune{'\u2660', '\u2661', '\u2662', '\u2663'}

func NewDeck() (deck Deck){
	for _,suit := range suits {
		for _,rank := range ranks {
			deck.Cards = append(deck.Cards, Card{rank,string(suit)})
		}
	}
	return deck
}

func (deck Deck) Shuffle() Deck{
	q1 := nodequeue.Queue[Card]{}
	q2 := nodequeue.Queue[Card]{}

	mismatch := -5 + rand.Intn(11)
	var i int
	for ; i < 26 + mismatch; i++ {
		q1.Insert(deck.Cards[i])
	}

	for ;i < 52;i++ {
		q2.Insert(deck.Cards[i])
	}

	deck = Deck{}
	for {
		if q1.Size() ==0 ||q2.Size() == 0 {
			break
		}

		deck.Cards = append(deck.Cards, q1.Remove())
		deck.Cards =  append(deck.Cards,q2.Remove())
	}

	if q1.Size() == 0 {
		for {
			if q2.Size() == 0 {
				break
			}
			deck.Cards =  append(deck.Cards,q2.Remove())
		}
	}

	if q2.Size() == 0 {
		for {
			if q1.Size() == 0 {
				break
			}
			deck.Cards = append(deck.Cards, q1.Remove())
		}
	}

	return deck
}


func main(){
	deck := NewDeck()
	fmt.Println("\n Original Deck: ", deck)
	// cut 5 times
	for index := 0; index < 5; index++ {
		deck = deck.Shuffle()
	}
	fmt.Println("\nShuffle deck: ", deck)
}