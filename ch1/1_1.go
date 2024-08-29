package main

import "fmt"

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func main() {
	students := []string{} //empty slice
	result := addStudent(students, "Micheal")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	fmt.Print(result)
}