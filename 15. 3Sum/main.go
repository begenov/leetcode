package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("threeSum(nums): %v\n", threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for left := 0; left < len(nums)-2; left++ {
		// Skip all duplicates from left after first element
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		center := left + 1
		right := len(nums) - 1

		for center < right {

			sum := nums[center] + nums[right] + nums[left]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[left], nums[center], nums[right]})

				right--

				// Skip all duplicates from right
				for center < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				right--
			} else {
				// Increment num2Idx to increase sum value
				center++
			}
		}
	}
	return result

}
