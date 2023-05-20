package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "Hello how are you Contestant"
	var k = 4
	truncateSentence(s, k)

}
func truncateSentence(s string, k int) string {
	sliceStr := strToSliceStr(s)
	res := strings.Join(sliceStr[0:k], " ")

	return res
}

func strToSliceStr(s string) []string {
	var res string
	var sliceStr []string
	for _, v := range s {
		if v == ' ' {
			fmt.Println("res")
			sliceStr = append(sliceStr, res)
			res = ""
		} else {
			res += string(v)
		}
	}
	if res != "" {
		sliceStr = append(sliceStr, res)
	}
	return sliceStr
}
