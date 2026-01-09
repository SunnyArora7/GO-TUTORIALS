package main

import (
	"fmt"
	"slices"
)

func main() {
	//Syntax

	var nameSlice []string
	nameSlice2 := []string{"sunny", "raj"}
	fmt.Println(nameSlice2)
	fmt.Println(nameSlice)
	fmt.Println(nameSlice == nil)

	ageSlice := make([]int, 0, 6)
	ageSlice = append(ageSlice, 9)
	fmt.Println(len(ageSlice))
	fmt.Println(ageSlice)
	fmt.Println(cap(ageSlice))

	cpyAgeSlice := make([]int, len(ageSlice))
	copy(cpyAgeSlice, ageSlice)
	fmt.Println(ageSlice, cpyAgeSlice)

	fmt.Println(slices.Equal(cpyAgeSlice, ageSlice))

	cpyAgeSlice = append(cpyAgeSlice, 5)
	fmt.Println(cpyAgeSlice)
	fmt.Println(cpyAgeSlice[0:])

}
