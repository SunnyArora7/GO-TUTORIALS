package main

import "fmt"

func main() {
	// we generally use arrays when we know the size in advance , for dynamics size we use slices

	//Syntax
	var nameArrays [4]string

	//To initialize the values of arrays in one line
	ageArray := [4]int{1, 2, 3, 4}
	nameArrays[2] = "sunny"
	fmt.Println(nameArrays)
	fmt.Println(len(nameArrays))
	fmt.Println(ageArray)
}
