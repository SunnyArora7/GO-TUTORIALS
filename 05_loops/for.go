package main

import "fmt"

func main() {

	//for loop example
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//while loop using for loop
	var i int = 0
	for i < 3 {
		fmt.Println("hello", i)
		i++
	}

	///for loop using range latest
	for j := range 3 {
		fmt.Println(j)
	}

}
