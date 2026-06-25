package main

import "fmt"

func printSlice[T int | string, V string](items []T, value V) {
	for _, item := range items {
		fmt.Println(item, value)
	}
}

// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }
func main() {
	nums := []int{1, 2, 3, 4, 5}
	// names := []string{"aditya", "g"}
	printSlice(nums, "adi")
	// printSlice(names)
	// printStringSlice(names)
}
