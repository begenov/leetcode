package main

import (
	"fmt"
	"sort"
)

func main() {
	piles := []int{2, 4, 1, 2, 7, 8, 21, 124}
	fmt.Println(maxCoins(piles))

}
func maxCoins(piles []int) int {

	sort.Ints(piles)
	out := 0
	fmt.Println(piles)
	first := 0
	last := len(piles) - 2
	fmt.Println(last, len(piles))
	for first < last {
		out += piles[last]
		first++
		last -= 2
	}

	return out
}
