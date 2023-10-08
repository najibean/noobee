package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s Student) CallMyName() {
	fmt.Println("Hello, My Name is " + s.Name)
}

func main() {
	student := Student{Name: "Najib", Class: "Intermediate"}
	student.CallMyName()

	student.SetMyName("Reihan")
	student.CallMyName()
}
