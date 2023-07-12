package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}

	fmt.Printf("singleNumber(nums): %v\n", singleNumber(nums))

}

// func singleNumber(nums []int) int {
// 	m := make(map[int]int, len(nums))
// 	for _, v := range nums {
// 		m[v]++
// 		if m[v] > 1 {
// 			delete(m, v)
// 		}

// 	}

// 	for v := range m {
// 		return v
// 	}

// 	return 0
// }

func singleNumber(nums []int) int {
	maps := make(map[int]struct{}, len(nums))

	for _, v := range nums {

		if _, ok := maps[v]; !ok {
			maps[v] = struct{}{}
		} else {
			delete(maps, v)
		}
	}

	for v := range maps {
		return v
	}
	return 0
}
