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

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

// func addStudentID(students []int, student int) []int {
// 	return append(students, student)
// }

// func addStudentStruct(students []Student, student Student) []Student {
// 	return append(students, student)
// }


// Group of functions that ensure that an OrderedSlice can be sorted

type Ordered interface {
	~int | ~float64 | ~string
}
 type OrderedSlice[T Ordered] []T //T myst implement < and >

 func (s OrderedSlice[T]) Len() int {
	return len(s)
 }

 func (s OrderedSlice[T]) Less(i,j int) bool {
	return  s[i] < s[j]
 }

 func (s OrderedSlice[T]) Swap(i,j int) {
	s[i], s[j] = s[j], s[i]
 }

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

// Map functions
func MyMap(input []int, f func(int) int) []int{
	result := make([]int,len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func GenericMap[T1,T2 any](input []T1,f func(T1)T2) []T2 {
	result := make([]T2, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func MyFilter(input []float64, f func (float64) bool) []float64 {
	var result []float64
	for _,value := range input {
		if f(value) {
			result = append(result,value)
		}
	}
	return result
}

func main() {
	students := []string{} //empty slice
	result := addStudent[string](students, "Micheal")
	result = addStudent[string](result, "Jennifer")
	result = addStudent[string](result, "Elaine")
	sort.Sort(OrderedSlice[string](result))
	fmt.Println(result)

	students1 := []int{} //empty slice
	results1 := addStudent[int](students1, 155)
	results1 = addStudent[int](results1, 122)
	results1 = addStudent[int](results1, 120)
	sort.Sort(OrderedSlice[int](results1))
	fmt.Println(results1)

	students2 := []Student{} //Empty slice
	results2 := addStudent[Student](students2, Student{"John", 213, 17.5})
	results2 = addStudent[Student](results2, Student{"James", 111, 18.75})
	results2 = addStudent[Student](results2, Student{"Marsha", 110, 16.25})
	PerformSort[Student](results2,func(s1, s2 Student) bool {
		return s1.Age <s2.Age // Compare twi Student values
	})
	fmt.Println(results2)

	// map functions
	slice := []int{1,5,2,7,4}
	result3 := MyMap(slice,func(i int) int {
		return i * i
	})
	fmt.Println(result3)

	// myfilter
	input := []float64{17.3, 11.1,9.9,4.3,12.6}
	res := MyFilter(input, func(f float64) bool {
		return f < 10.0 
	})
	fmt.Println(res)
}
