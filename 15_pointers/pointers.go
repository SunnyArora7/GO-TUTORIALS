package main

import "fmt"

// Call By Value
func callByValue(num int) {
	num = 5
	fmt.Println("In Method Value in Call By Value:", num)
}

// Call By Reference
func callByReference(num *int) {
	*num = 6
	fmt.Println("In Method Value in Call By Reference:", *num)
}
func main() {
	num := 1
	callByValue(num)
	fmt.Println("After Method Call by Value :", num)

	num2 := 2
	callByReference(&num2)
	fmt.Println("After Call By Reference num2 value", num2)

}
