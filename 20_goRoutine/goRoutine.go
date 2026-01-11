package main

import (
	"fmt"
	"sync"
)

// Go routines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions, enabling efficient multitasking within a Go program. To create a goroutine, you use the `go` keyword followed by a function call. This starts the function in a new goroutine, allowing it to run concurrently with other goroutines.

func displayNum(num int, wg *sync.WaitGroup) {
	//Here we use defer so the wg.done will be executed after the function is executed
	defer wg.Done()
	fmt.Println("Working on number:", num)
}
func main() {
	//syntax for wait group
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go displayNum(i, &wg)
	}

	wg.Wait()

	//time.Sleep(2 * time.Second) we generally avoid using sleep to wait for goroutines to finish
	//We will use waitgroups for it so that all the go routines complete before main function exits
}
