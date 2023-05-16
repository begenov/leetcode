package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{1, 1, 2}
	log.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	var uniqueNums []int

	for _, v := range nums {
		if uniqueElements(uniqueNums, v) {
			uniqueNums = append(uniqueNums, v)
		}
	}

	fmt.Println(uniqueNums)

	return len(uniqueNums)
}

func uniqueElements(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return false
		}
	}
	return true
}
