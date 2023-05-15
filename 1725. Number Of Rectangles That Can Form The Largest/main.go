package main

import "log"

func main() {
	count := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}

	log.Println(countGoodRectangles(count))
}

func countGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	count := 0
	for _, r := range rectangles {
		sideLen := min(r[0], r[1])
		if sideLen > maxLen {
			maxLen = sideLen
			count = 1
		} else if sideLen == maxLen {
			count++
		}
	}
	return count
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
