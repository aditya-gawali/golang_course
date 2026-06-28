package main

import (
	"fmt"
	"time"
)

// // send channel
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("processing num ", num)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum(result chan int, num1 int, num2 int){
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func task(done chan bool) {
// 	defer func() {done <- true}()
// 	fmt.Println("processing...")
// }

// send only and receive only channel
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending mail to ", email)
		time.Sleep(time.Second)
	}

}

func main() {
	// numChan := make(chan int)

	// go processNum(numChan) // async

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second)

	// result := make(chan int)

	// go sum(result, 3,4)

	// sum := <- result

	// fmt.Println(sum)

	// done := make(chan bool)
	// go task(done)
	// <- done

	// emailChan := make(chan string, 100) // buffered channel
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("sending done")
	// close(emailChan)
	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Data := <-chan1:
			fmt.Println(chan1Data)
		case chan2Data := <-chan2:
			fmt.Println(chan2Data)
		}
	}

}
