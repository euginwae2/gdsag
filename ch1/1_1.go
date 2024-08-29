package main

import "fmt"

type Student struct {
	Name string
	ID   int
	age  float64
}

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func addStudentID(students []int, student int) []int {
	return append(students, student)
}

func addStudentStruct(students []Student, student Student) []Student {
	return append(students, student)
}

func main() {
	students := []string{} //empty slice
	result := addStudent(students, "Micheal")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	fmt.Println(result)

	students1 := []int{} //empty slice
	results1 := addStudentID(students1, 155)
	results1 = addStudentID(results1, 122)
	results1 = addStudentID(results1, 120)

	fmt.Println(results1)

	students2 := []Student{} //Empty slice
	results2 := addStudentStruct(students2, Student{"John", 213, 17.5})
	results2 = addStudentStruct(results2, Student{"James", 111, 18.75})
	results2 = addStudentStruct(results2, Student{"Marsha", 110, 16.25})
	fmt.Println(results2)
}
