package main

import "fmt"

func main() {

	if age:= 11; age >= 18 {
		fmt.Println("you are adult", age)
	}else if age >= 12 {
		fmt.Println("you are teenager", age)
	}else{
		fmt.Println("you are a kid", age)
	}

	// not support for ternery operator
}