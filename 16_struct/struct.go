package main

import (
	"fmt"
	"time"
)

//  struct embedding
type customer struct{
	name string
	phone string
}

type order struct {
	id       string
	amount   float64
	status   string
	createAt time.Time
	customer
}

// receiver type -  bind the func to struct function binding 

func (o *order) changeStatus (status string){
	o.status = status
}

func (o order) getAmount () float64{
	return o.amount
}

// struct constructor 
func newOrder(id string, amount float64, status string) *order{
	myOrder := order{
		id: id,
		amount : amount,
		status : status,
	}

	return &myOrder

	
}

func main() {
	// myOrder := order{
	// 	id : "123", amount: 50, status: "received",
	// }
	// myOrder2 := order{
	// 	id : "234", amount: 100, status: "paid",
	// }

	// myOrder.createAt = time.Now()

	// fmt.Println(myOrder)
	// fmt.Println(myOrder2)
	// myOrder2.changeStatus("delivered")
	// fmt.Println(myOrder2.status)
	// fmt.Println(myOrder2.getAmount())

	myOrder1 := newOrder("1", 100, "paid")

	fmt.Println(myOrder1)


	//inline struct 
	languages := struct {
		name string
		isGood bool
	}{"golang", true}

	fmt.Println(languages)

	newCustomer := customer{
		name : "aditya",
		phone: "1234567890",
	}
	// struct embedding
	myOrder := order{
		id: "1234",
		amount: 123,
		status: "received",
		createAt: time.Now(),
		customer: newCustomer,
	}

	fmt.Println(myOrder.name)

}