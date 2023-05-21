package main

import "log"

func main() {
	var n = []int{1, 2, 3, 3}
	log.Println(repeatedNTimes(n))
}
func repeatedNTimes(nums []int) int {
	var m = make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	maxCount := 0
	k := 0
	for key, value := range m {
		if value > maxCount {
			maxCount = value
			k = key
		}
	}

	return k

}
