package main

import "fmt"

// Generally we create a interface with the name ends with er
// In go interface not needs to be called explicitly it will be implicitly call if any structs uses the same name method with same parameters
type paymenter interface {
	pay(amount float32)
}
type paymentWith struct {
	payType paymenter
}

func (p paymentWith) makePayment(amount float32) {
	p.payType.pay(amount)
}

type stripe struct {
}

type razorPay struct {
}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment", amount, "with stripe")
}

func (r razorPay) pay(amount float32) {
	fmt.Println("Making payment", amount, "with razorpay")
}

func main() {
	razaorPayOrdr := razorPay{}
	stripeOrder := stripe{}

	paymentWithOrder1 := paymentWith{payType: razaorPayOrdr}
	paymentWithOrder2 := paymentWith{payType: stripeOrder}

	paymentWithOrder1.makePayment(3000)
	paymentWithOrder2.makePayment(3000)

}
