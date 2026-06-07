package main

func main() {
	// slice -> dynamic array

	// var nums []int

	// fmt.Println(nums)

	// make func to create slice

	// var nums = make([]int, 0, 5)

	// fmt.Println(nums)

	// fmt.Println(len(nums)) // check the length of the slice
	// fmt.Println(cap(nums))  // check the capacity of the slice

	// nums = append(nums, 1)

	// fmt.Println(nums)

	// another way to create slice
	// nums := []int{}

	// nums = append(nums, 1,2,3,4,5)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	//copy func

	// var nums = make([]int,0,5)
	// nums = append(nums, 2)

	// var num2 = make([]int, len(nums))
	// fmt.Println(nums,num2)

	// copy(num2,nums)

	// fmt.Println(nums,num2)

	//slice operator

	// var nums = []int{1,2,3}

	// fmt.Println(nums[0:2])  // slice operator

	// fmt.Println(nums[:2])

	// fmt.Println(nums[1:])

	// var nums1 = []int{1,2}
	// var nums2 = []int{1,2}
	// // slices pkg
	// fmt.Println(slices.Equal(nums1, nums2))

	//2d slices

	// var nums = [][]int{{1,2},{3,4}}

	// nums = append(nums, []int{1, 2})

	// nums[0] = append(nums[0], 3, 4)

	// fmt.Println(nums)

}
