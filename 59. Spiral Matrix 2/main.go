package main

import (
	"log"
)

func main() {
	log.Println(generateMatrix(4))
}

// [[1,2],[4,3]]

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	rowStart, rowEnd := 0, n-1
	colStart, colEnd := 0, n-1

	for rowStart <= rowEnd && colStart <= colEnd {

		for i := colStart; i <= colEnd; i++ {
			matrix[rowStart][i] = num
			num++
		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			matrix[i][colEnd] = num

			num++
		}
		colEnd--

		if rowStart <= rowEnd {

			for i := colEnd; i >= colStart; i-- {
				matrix[rowEnd][i] = num
				num++
			}
			rowEnd--
		}

		if colStart <= colEnd {

			for i := rowEnd; i >= rowStart; i-- {
				matrix[i][colStart] = num
				num++
			}
			colStart++

		}
	}

	return matrix
}

/*
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
*/
