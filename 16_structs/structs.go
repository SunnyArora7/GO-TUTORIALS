package main

import "fmt"

type student struct {
	name       string
	age        int
	medicalFit bool
}

//In oops we create constructor to initialize the object , in go we create method just below the struct

func newStudent(name string, age int, medicalFit bool) *student {
	myStudent := student{name: name, age: age, medicalFit: medicalFit}
	return &myStudent
}

// Here before the function name we just create a receiver to connect with struct
func (s *student) changeName(name string) {
	s.name = name
}
func (s student) getName() string {
	return s.name
}

func main() {
	sunny := student{name: "sunny", age: 24, medicalFit: true}
	raj := student{name: "raj", age: 25, medicalFit: false}
	bhavik := newStudent("Bhavik Chawla", 24, false)
	sunny.age = 28
	fmt.Println(sunny)
	fmt.Println(raj)
	fmt.Println(sunny.name)

	sunny.changeName("Sunny Arora")

	fmt.Println(sunny.getName())
	fmt.Println(bhavik)

	//****** Inline Struct********
	language := struct {
		name   string
		isGood bool
	}{"java", false}
	fmt.Println(language)
}
