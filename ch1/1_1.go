package main

import "fmt"

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func addStudentID(students []int, student int) []int{
	return append(students,student)
}

func main() {
	students := []string{} //empty slice
	result := addStudent(students, "Micheal")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	fmt.Println(result)

	students1 := []int{} //empty slice
	results1 := addStudentID(students1,155)
	results1 = addStudentID(results1,122)
	results1 = addStudentID(results1,120)

	fmt.Println(results1)
	
	
}