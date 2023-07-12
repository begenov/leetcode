package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1} // 1 2 3 4 ... ... 7 8

	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {

	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}
	min := 1
	res := []int{}
	for ; min <= len(nums); min++ {
		if _, ok := m[min]; !ok {
			res = append(res, min)
		}
	}

	return res
}
