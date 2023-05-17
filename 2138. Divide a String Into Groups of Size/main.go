package main

import (
	"fmt"
)

func main() {
	s := "abcdefghi1"
	k := 3
	fill := byte('x')
	fmt.Println(divideString(s, k, fill))

}

func divideString(s string, k int, fill byte) []string {
	var res []string
	var word string
	for _, x := range s {
		word += string(x)
		if len(word) == k {
			res = append(res, word)
			word = ""
		}
	}

	if word != "" {
		word = addInSliceByte(word, fill, k)
		res = append(res, word)
	}

	return res

}

func addInSliceByte(x string, fill byte, n int) string {
	for i := 0; i < n; i++ {
		if len(x) == n {
			return x
		}
		x += string(fill)
	}
	return x
}
