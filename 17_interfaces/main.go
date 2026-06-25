package main

import "fmt"

type paymenter interface{
	pay(amount float32)
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// // razorpayPaymentGW := razorpay{}
	// // razorpayPaymentGW.pay(amount)

	// stripePaymentGW := stripe{}
	// stripePaymentGW.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make a payment

	fmt.Println("making payment using razorpay for amount = ", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("payking payment using stripe for amount = ", amount)
}

type fakePayment struct{}

func (s fakePayment) pay(amount float32){
	fmt.Println("payking payment using fake for testing purpose = ", amount)
}

func main() {
	// razorpayGW := razorpay{}
	stripeGW := stripe{}
	newPayment := payment{
		gateway : stripeGW,
	}
	newPayment.makePayment(100)
}