package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollNo  int
	name    string
	address string
}

type StudentList struct {
	students []Student
}

func newStudent(rollNo int, name string, address string) *Student {
	return &Student{rollNo, name, address} // because go provides automatic memory management (it will automatically move the object to heap)
}

func (studentList *StudentList) CreateStudent(rollNo int, name string, address string) {
	student := newStudent(rollNo, name, address)
	studentList.students = append(studentList.students, *student)
}

func (studentList *StudentList) printStudents() {
	for index, student := range studentList.students {
		fmt.Println(strings.Repeat("=", 15), "Student", index+1, strings.Repeat("=", 15))
		fmt.Println("Roll No:", student.rollNo)
		fmt.Println("Name:", student.name)
		fmt.Println("Address:", student.address)
	}
}

func main() {
	studentList := new(StudentList)

	studentList.CreateStudent(1, "Devil", "Hell")
	studentList.CreateStudent(2, "Angel", "Heaven")
	studentList.CreateStudent(3, "Human", "Earth")

	studentList.printStudents()
}
