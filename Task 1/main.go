package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(p Person) {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("Job: ", p.job)
	fmt.Println("Salary: ", p.salary)
	fmt.Println()
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	fmt.Println("Person 1")
	printPerson(pers1)

	fmt.Println("Person 2")
	printPerson(pers2)
}
