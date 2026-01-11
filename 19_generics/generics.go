package main

import "fmt"

func display[T any](data []T) {
	for index, item := range data {
		fmt.Println(item, index)
	}
}

type genericsStruct[T int | float32] struct {
	data []T
}

func main() {
	//nums := []int{12, 24, 32, 42, 41}
	floatNums := []float32{11.44, 1123.3}
	display(floatNums)

	display([]string{"sunny", "raj"})

	student1 := genericsStruct[int]{data: []int{12, 123}}
	fmt.Println(student1)
}
