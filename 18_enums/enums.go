package main

import "fmt"

type orderStatus string

const (
	//packed orderStatus = iota this is for the int type
	packed     orderStatus = "packed"
	dispatched orderStatus = "dispatched"
	delivered  orderStatus = "delivered"
)

func checkDeliveryStatus(status orderStatus) {
	fmt.Println("Delivery Status:", status)
}
func main() {
	checkDeliveryStatus(delivered)
}
