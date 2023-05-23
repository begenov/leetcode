package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 0, 1}
	fmt.Printf("missingNumber(nums): %v\n", missingNumber(nums))
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}
	i, j := 0, 0
	for ; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			j = nums[i+1] - 1
			break
		}
	}
	if j == 0 {
		return len(nums)
	}
	return j
}
