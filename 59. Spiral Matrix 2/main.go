package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	var res [][]int
	max := n * 3

	for {
		var row []int
		for i := 1; i <= n; i++ {
			row = append(row, i)
		}

		res = append(res, row)
		var row2 []int
		var row3 []int
		c1 := 0

		count := 2
		for i := n + 1; i <= max; i++ {
			if i == n+1 {
				c1 = i
				continue
			}
			if i == n+2 {
				row3 = append(row3, i)
				continue
			}
			if count <= 0 {
				row2 = append(row2, i)
			} else {
				row3 = append(row3, i)
			}
			count -= 1
			fmt.Println(count)
		}
		row2 = append(row2, c1)

		res = append(res, row2)
		var r []int
		for i := len(row3) - 1; i >= 0; i-- {
			r = append(r, row3[i])
		}
		res = append(res, r)
		break
	}
	return res
}
