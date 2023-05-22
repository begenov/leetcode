package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Println("oo")
	merge(nums1, m, nums2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < m {
		m = len(nums1)
	}
	if len(nums2) < n {
		n = len(nums2)
	}
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	nums1 = add(nums1, nums2)
	sort.Ints(nums1)
}

func add(nums1, nums2 []int) []int {
	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}
	return nums1
}
