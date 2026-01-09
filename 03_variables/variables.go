package main

import "fmt"

func main() {
	fmt.Println("Variables in go")

	// this is used when the value will be declared at run time
	var name string = "sunny"
	fmt.Println(name)

	// this syntax implicitly takes the type of variable when we set the value
	var name2 = "raj"
	fmt.Println(name2)

	//Short hand syntax used when value is initialized while declaring
	name3 := "bhavik"
	fmt.Println(name3)

	var age int = 32
	fmt.Println(age)

	var height float32 = 167.6
	fmt.Println(height)

}
