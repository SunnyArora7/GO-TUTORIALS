package main

import (
	"fmt"
	"time"
)

func main() {
	//Syntax
	age := 18
	switch age {
	case 18:
		fmt.Println("Adult")
	default:
		fmt.Println("Not an adult")
	}

	//condition based switch statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekdays")
	}

	//Type switch
	check := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("its integer")
		case float32, float64:
			fmt.Println("its float")

		default:
			fmt.Println("Its bool")
		}
	}

	check(false)
}
