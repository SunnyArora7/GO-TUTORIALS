package main

import "fmt"

func closureFunc() func() int {
	var total int
	return func() int {
		//In this closure concepts comes if this is a normal function then the total value b deleted after every function close
		//But in this the outer variable will be stay
		total = total + 1
		return total
	}
}
func main() {
	increment := closureFunc()
	fmt.Println(increment())
	fmt.Println(increment())

}
