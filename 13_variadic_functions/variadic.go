package main

import "fmt"

//variadic functions is a function which can accept multiple values of any types

func variadicFunc(notepad ...any) any {
	return notepad
}

func variadicFunc2(num ...int) int {
	var sum int
	for i := range num {
		sum = sum + num[i]
	}
	return sum
}
func main() {
	fmt.Println(variadicFunc(1, 2, 4, 4, "string"))

	fmt.Println(variadicFunc2(1, 23, 44, 4))

	//we can also pass like this
	slices := []int{1, 2, 3}
	fmt.Println(variadicFunc2(slices...))

}
