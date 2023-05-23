package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("majorityElement(nums): %v\n", majorityElement(nums))

}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	max := 0
	k := 0
	for key, value := range m {
		if value > max {
			max = value
			k = key
		}
	}
	return k
}
