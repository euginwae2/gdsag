package main

import (
	"fmt"

	"example.com/nodequeue"
)


type Passenger struct {
	name string
	priority int
}

type PriorityQueue[T any] struct {
	q    []nodequeue.Queue[T] //slice of queue
	size int
}

func NewPriorityQueue[T any](numberOfPriorities int) (pq PriorityQueue[T]) {
	pq.q = make([]nodequeue.Queue[T], numberOfPriorities)
	return pq
}

// Methods for priority queue
// Insert
func (pq *PriorityQueue[T]) Insert(item T, priority int) {
	pq.q[priority -1].Insert(item)
	pq.size++
}

// Remove
func (pq *PriorityQueue[T]) Remove() T {
	pq.size--
	for i := 0; i < len(pq.q); i++ {
		if pq.q[i].Size() > 0 {
			return pq.q[i].Remove()
		}
	}
	var zero T
	return zero
}

// First
func (pq *PriorityQueue[T]) First() T{
	for _, queue := range(pq.q){
		if queue.Size() > 0 {
			return queue.First()
		}
	}
	var zero T
	return zero
}

// IsEmpty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	for i :=0; i < len(pq.q); i++ {
		if pq.q[i].Size() > 0 {
			return false
		}
	}
	return true
}

func main() {
	airlineQueue := NewPriorityQueue[Passenger](3)
	passegers :=  []Passenger{
		{"Erika",3}, {"Robert", 3}, {"Danielle", 3}, {"Madison",1}, {"Fredrick",1}, {"James",2},
		{"Dante", 2}, {"Shelley",3},
	}

	fmt.Println("Passenger: ", passegers)
	for i :=0; i < len(passegers); i++ {
		airlineQueue.Insert(passegers[i], passegers[i].priority)
	}
	fmt.Println("First passenmger in line: ", airlineQueue.First())
	airlineQueue.Remove()
	airlineQueue.Remove()
	airlineQueue.Remove()
	fmt.Println("First passenger in line ", airlineQueue.First())
}