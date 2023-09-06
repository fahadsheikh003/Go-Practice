package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	position string
}

type Company struct {
	companyName string
	employees   []Employee
}

func main() {
	e1 := Employee{"John", 5000, "Manager"}
	e2 := Employee{"Mary", 6000, "Senior Manager"}
	e3 := Employee{"Mike", 4000, "Junior Manager"}

	employees := []Employee{e1, e2, e3}

	logic := Company{"Logic", employees}

	fmt.Println(logic)
}
