package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))

	for _, i := range nums {
		m[i] = struct{}{}
	}

	return len(m) != len(nums)
}
