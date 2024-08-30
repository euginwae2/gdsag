package main

import (
	"fmt"
	"sort"
)

type Stringer = interface {
	String() string
}

type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) String() string {
	return string(s)
}

type Student struct {
	Name string
	ID   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}

func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}

// func addStudentID(students []int, student int) []int {
// 	return append(students, student)
// }

// func addStudentStruct(students []Student, student Student) []Student {
// 	return append(students, student)
// }


// Group of functions that ensure that an OrderedSlice can be sorted

//  type OrderedSlice[T Ordered] []T //T myst implement < and >

//  func (s OrderedSlice[T]) Len() int {
// 	return len(s)
//  }

//  func (s OrderedSlice[T]) Less(i,j int) bool {
// 	return  s[i] < s[j]
//  }

//  func (s OrderedSlice[T]) Swap(i,j int) {
// 	s[i], s[j] = s[j], s[i]
//  }

//  Group of functions that ensure that SortType can be sorted
 type SortType[T any] struct{
	slice []T
	compare func(T,T) bool
 }

 func (s SortType[T]) Len() int {
	return len(s.slice)
 }

 func (s SortType[T]) Less(i,j int) bool {
	return s.compare(s.slice[i], s.slice[j])
 }

 func (s SortType[T]) Swap(i,j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
 }

//  end group for SortType

func PerformSort[T any](slice []T,compare func(T,T) bool) {
	sort.Sort(SortType[T]{slice,compare})
}

func main() {
	students := []String{} //empty slice
	result := addStudent[String](students, "Micheal")
	result = addStudent[String](result, "Jennifer")
	result = addStudent[String](result, "Elaine")
	fmt.Println(result)

	students1 := []Integer{} //empty slice
	results1 := addStudent[Integer](students1, 155)
	results1 = addStudent[Integer](results1, 122)
	results1 = addStudent[Integer](results1, 120)

	fmt.Println(results1)

	students2 := []Student{} //Empty slice
	results2 := addStudent[Student](students2, Student{"John", 213, 17.5})
	results2 = addStudent[Student](results2, Student{"James", 111, 18.75})
	results2 = addStudent[Student](results2, Student{"Marsha", 110, 16.25})
	fmt.Println(results2)
}
