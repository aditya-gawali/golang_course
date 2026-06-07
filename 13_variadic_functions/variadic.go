package main

import "fmt"

func sum(nums ...int)int{
	total := 0

	for _, num := range nums{
		total += num
	}
	return total
}

func main() {

	// a func accept any no. of parameters

	fmt.Println(1,2,3,4,5)

	totalSum := sum(1,2,3,4,5,6,7)
	fmt.Println(totalSum)

	nums := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(sum(nums...))
}