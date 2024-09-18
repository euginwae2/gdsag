package main

import (
	"fmt"

	// "example.com/nodequeue"
	// "example.com/slicequeue"
	"example.com/deque"
)

// const size = 1_000_000

func main() {
	// myQueue := slicequeue.Queue[int]{}
	// myQueue.Insert(15)
	// myQueue.Insert(20)
	// myQueue.Insert(30)
	// myQueue.Remove()
	// fmt.Println(myQueue.First())
	// queue := slicequeue.Queue[float64]{}
	// for i := 0; i < 10; i++ {
	// 	queue.Insert(float64(i))
	// }
	// iterator := queue.Range()
	// for {
	// 	if iterator.Empty() {
	// 		break
	// 	}
	// 	fmt.Println(iterator.Next())
	// }
	// fmt.Println("Queue.First() = ", queue.First())

	// sliceQueue := slicequeue.Queue[int]{}
	// nodeQueue := nodequeue.Queue[int]{}
	// start := time.Now()
	// for i := 0; i < size; i++ {
	// 	sliceQueue.Insert(i)
	// }
	// elasped := time.Since(start)
	// fmt.Println("Time for inserting 1 million ints in sliceQueue is ", elasped)

	// start =  time.Now()
	// for i :=0; i<size; i++{
	// 	nodeQueue.Insert(i)
	// }
	// elasped = time.Since(start)
	// fmt.Println("Time for inserting 1 million ints in nodeQueue is ", elasped)

	// start = time.Now()
	// for i := 0; i < size; i++ {
	// 	sliceQueue.Remove()
	// }
	// elasped = time.Since(start)
	// fmt.Println("Time for removing 1 million ints in sliceQueue is ", elasped)

	// start =  time.Now()
	// for i :=0; i < size; i++ {
	// 	nodeQueue.Remove()
	// }
	// elasped =  time.Since(start)
	// fmt.Println("Time for remove 1 million ints in nodeQueue is ", elasped)

	myDeque := deque.Deque[int]{}
	myDeque.InsertFront(5)
	myDeque.InsertBack(10)
	myDeque.InsertFront(2)
	myDeque.InsertBack(12)
	fmt.Println("myDeque.First()= ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())

	myDeque.RemoveLast()
	myDeque.RemoveFirst()
	fmt.Println("myDeque.First() = ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())
}
