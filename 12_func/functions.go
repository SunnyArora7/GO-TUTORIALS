package main

import "fmt"

// Basic syntax for function
func add(a, b int) int {
	return a + b
}

//We can return multiple values from the functions special feature of go

func returnMul() (string, int, bool) {
	return "sunny", 10, true
}

// we can return functions from the functions like the below example
func funReturnFunc() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main() {
	//_ is used when we dont want to use the variables so instead of bool variable we have used_
	name, age, _ := returnMul()
	fmt.Println(add(10, 20))
	fmt.Println(name, age)

	//here we are assigning a function to a variable
	f := funReturnFunc()
	fmt.Println(f(20, 20))
}
