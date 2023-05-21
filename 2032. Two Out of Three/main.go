package main

import "fmt"

func main() {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}

	result := twoOutOfThree(nums1, nums2, nums3)
	fmt.Println(result) // Output: [3 2]

	nums1 = []int{3, 1}
	nums2 = []int{2, 3}
	nums3 = []int{1, 2}

	result = twoOutOfThree(nums1, nums2, nums3)
	fmt.Println(result) // Output: [2 3 1]

	nums1 = []int{1, 2, 2}
	nums2 = []int{4, 3, 3}
	nums3 = []int{5}

	result = twoOutOfThree(nums1, nums2, nums3)
	fmt.Println(result)
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	m3 := make(map[int]struct{})

	for _, num := range nums1 {
		m1[num] = struct{}{}
	}
	for _, num := range nums2 {
		m2[num] = struct{}{}
	}
	for _, num := range nums3 {
		m3[num] = struct{}{}
	}

	var ret []int
	for i := 1; i <= 100; i++ {
		_, ok1 := m1[i]
		_, ok2 := m2[i]
		_, ok3 := m3[i]
		if (ok1 && ok2) || (ok2 && ok3) || (ok3 && ok1) {
			ret = append(ret, i)
		}
	}
	return ret

}
