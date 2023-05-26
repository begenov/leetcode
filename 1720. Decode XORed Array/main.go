package main

import "fmt"

func main() {
	encoded := []int{6, 2, 7, 3}
	first := 4
	fmt.Printf("decode(encoded, first): %v\n", decode(encoded, first))

}

func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	fmt.Println(ans)

	for idx, sum := range encoded {
		ans[idx+1] = sum ^ ans[idx]
		fmt.Println(ans)
	}
	return ans
}
