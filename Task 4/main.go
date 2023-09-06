package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	rollNo  int
	name    string
	address string
	courses []string
}

type StudentList struct {
	students []Student
}

func newStudent(rollNo int, name string, address string, courses []string) *Student {
	return &Student{rollNo, name, address, courses} // because go provides automatic memory management (it will automatically move the object to heap)
}

func (studentList *StudentList) CreateStudent(rollNo int, name string, address string, courses []string) {
	student := newStudent(rollNo, name, address, courses)
	studentList.students = append(studentList.students, *student)
}

func concatenateStrings(seperator string, strings ...string) string {
	var result string
	for index, str := range strings {
		if index == len(strings)-1 {
			result += str
		} else {
			result += str + seperator
		}
	}
	return result
}

func (student *Student) getHash() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(concatenateStrings("", strconv.Itoa(student.rollNo), student.name, student.address, concatenateStrings("", student.courses...)))))
}

func (studentList *StudentList) printStudents() {
	for index, student := range studentList.students {
		fmt.Println(strings.Repeat("=", 15), "Student", index+1, strings.Repeat("=", 15))
		fmt.Println("Roll No:", student.rollNo)
		fmt.Println("Name:", student.name)
		fmt.Println("Address:", student.address)
		fmt.Println("Courses:", concatenateStrings(", ", student.courses...))
		fmt.Println("Hash:", student.getHash())
	}
}

func main() {
	studentList := new(StudentList)

	studentList.CreateStudent(1, "Devil", "Hell", []string{"Maths", "Science"})
	studentList.CreateStudent(2, "Angel", "Heaven", []string{"Maths", "Science", "English"})
	studentList.CreateStudent(3, "Human", "Earth", []string{"Maths", "Science", "English", "German"})

	studentList.printStudents()
}
