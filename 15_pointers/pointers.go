package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In ChangeNum", *num)
}

func main() {
	num := 1

	fmt.Println("memory address", &num)
	changeNum(&num)

	fmt.Println("after changeNum in main", num)

}