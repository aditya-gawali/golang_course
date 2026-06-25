package main

import "fmt"

type orderStatus int

const (
	Received orderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("The status is changed to ", status)
}

func main() {
	changeOrderStatus(Received)
}