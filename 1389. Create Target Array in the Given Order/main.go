package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	fmt.Printf("createTargetArray(nums, index): %v\n", createTargetArray(nums, index))
}

func createTargetArray(nums []int, index []int) []int {
	var target []int

	for i := 0; i < len(index); i++ {
		t := nums[i]
		ii := index[i]

		target = append(target[:ii], append([]int{t}, target[ii:]...)...)

	}
	return target
}

// func createTargetArray(nums []int, index []int) []int {
// 	target := make([]int, len(nums))

// 	for i := 0; i < len(index); i++ {
// 		idx := index[i]
// 		num := nums[i]

// 		copy(target[idx+1:], target[idx:])
// 		fmt.Println(target)
// 		target[idx] = num
// 		fmt.Println(target)

// 	}

// 	return target
// }
