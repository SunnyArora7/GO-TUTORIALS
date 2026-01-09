package main

import "fmt"

func main() {

	//Normal syntax
	var age int = 18
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not an adult")
	}

	//We can declare a variable in if statement
	if graduation := true; graduation == true {
		fmt.Println("Graduated")
	} else {
		fmt.Println("Not graduated")
	}

	//go does not have ternary operator
}
