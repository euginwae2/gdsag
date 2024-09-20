package main

import (
	"fmt"

	"example.com/set"
)

func main() {
	set1 := set.Set[int]{}
	set1.Insert(3)
	set1.Insert(5)
	set1.Insert(7)
	set1.Insert(9)

	set2 := set.Set[int]{}
	set2.Insert(3)
	set2.Insert(6)
	set2.Insert(8)
	set2.Insert(9)
	set2.Insert(11)
	set2.Delete(11)

	fmt.Println("Items in set2: ", set2.Items())
	fmt.Println("5 is in set1: ", set1.In(5))
	fmt.Println("5 is in set2: ", set2.In(5))

	fmt.Println("union of set1 and set2: ", set1.Union(set2).Items())
	fmt.Println("intersection of set1 and set2: ", set1.Intersection(set2).Items())
	fmt.Println("Difference of set2 with respect to set1: ", set2.Difference(set1).Items())
	fmt.Println("Size of this difference: ", set1.Intersection(set2).Size())
}
