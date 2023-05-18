package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func main() {
	input := []int{10, 20, 30, 5, 10, 50}
	log.Println(maxAscendingSum(input))
	f := make(map[int]int)
	fmt.Println(f[2])
	context.TODO()
	context.Background()
	sync.Map
}

func maxAscendingSum(nums []int) int {
	sum1, sum2 := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		sum1 = max(sum1, nums[i])

		if nums[i] > nums[i-1] {
			sum2 += nums[i]
		} else {
			sum1 = max(sum1, sum2)
			sum2 = nums[i]
		}
	}

	return max(sum1, sum2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
