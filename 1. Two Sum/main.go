package main

import "log"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	log.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var a, b int
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			a = nums[i]
			b = nums[j]
			if a+b == target {
				res = append(res, i)
				res = append(res, j)
				return res
			}
		}
	}
	return res
}
