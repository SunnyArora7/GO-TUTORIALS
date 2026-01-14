package main

import (
	"fmt"
	"sync"
)

// Receiving method example of channel
// in this method we can only receive the data as mentioned in the paramters <-chan its state we can only receive data not send
func sendData(messageChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Data Sended from main go routine", <-messageChan)
}

// send emthod exampleof channel
func receiveData(messageChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messageChan <- "hello"
}

func main() {
	//Method Channels are basically used to send data btw go routines
	//Syntax to declare a channel
	var wg sync.WaitGroup
	wg.Add(1)
	messageChan := make(chan string)
	go receiveData(messageChan, &wg)
	receiveDataText := <-messageChan
	fmt.Println(receiveDataText)
	//go sendData(messageChan, &wg)
	//messageChan <- "Ping" //Here are we are sending the message into the channel
	wg.Wait()

	close(messageChan)
	//<-messageChan //Here we are receiving the message from the channel

}
