package main

import "fmt"

func main() {
	var s = []string{"d", "b", "c", "b", "c", "a"}
	var k = 2
	fmt.Println(kthDistinct(s, k))

}
func kthDistinct(arr []string, k int) string {
	var i int
	var b bool
	var res []string

	for ; i < len(arr); i++ {
		b = repeat(arr, arr[i], i)
		if !b {
			continue
		} else {
			res = append(res, arr[i])
		}
	}

	if len(res) < k {
		return ""
	}
	return res[k-1]
}

func repeat(str []string, s string, in int) bool {

	for i := 0; i < len(str); i++ {
		if i == in {
			continue
		}
		if str[i] == s {
			return false
		}
	}
	return true
}
