package main

import (
	"fmt"
)

func main() {
	s := "bgycymgbblobvpdf"
	k := 67
	fill := byte('u')
	fmt.Println(divideString(s, k, fill))

}

func divideString(s string, k int, fill byte) []string {
	var res []string
	var word string
	count := 0
	for _, v := range s {

		if count < k {
			word += string(v)
			count++
		} else {
			res = append(res, word)
			word = ""
			count = 0
		}

	}

	return res

}
