package main

import "fmt"

func main() {
	//This is the example for slices iteration

	var nums = []int{1, 2, 3}

	//First syntax
	for i := range nums {
		fmt.Println(i, nums[i])
	}
	fmt.Println("*******************")
	//second syntax
	for i, num := range nums {
		fmt.Println(i, num)
	}

	//This is the example for maps
	m := make(map[string]string)
	m["name"] = "sunny"

	//in this i is the key , and e is the value of the key
	for i, e := range m {
		fmt.Println(i, e)
	}

	//This is the example on string
	for i, c := range "sunny" {
		fmt.Println(i, string(c))
	}

}
