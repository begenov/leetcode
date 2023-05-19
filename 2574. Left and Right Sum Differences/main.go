package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 4, 8, 3}
	fmt.Println(leftRightDifference(nums))
}
func leftRightDifference(nums []int) []int {
	var left = make([]int, 0, len(nums))
	left = append(left, 0)
	// var right = make([]int, 0, len(nums))
	for i := 1; i < len(nums); i++ {
		left = append(left, left[i-1]+nums[i-1])

	}

	right := make([]int, 0, len(nums))

	right = appendSlice(len(nums))

	for i := len(nums) - 2; i >= 0; i-- {
		fmt.Println(i)

		right[i] = right[i+1] + nums[i+1]
	}

	var res = make([]int, 0, len(nums))

	res = appendSlice(len(nums))

	for i := 0; i < len(nums); i++ {
		if left[i]-right[i] >= 0 {
			res[i] = (left[i] - right[i]) * 1
		} else {
			res[i] = (left[i] - right[i]) * -1

		}

	}

	return res
}

func appendSlice(nums int) []int {
	var newSlice []int
	for i := 0; i < nums; i++ {
		newSlice = append(newSlice, 0)
	}
	return newSlice
}
