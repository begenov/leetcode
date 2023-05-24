package main

import "fmt"

func main() {
	fmt.Printf("alternateDigitSum(555): %v\n", alternateDigitSum(101))
}
func alternateDigitSum(n int) int {
	var result int
	for ; n != 0; n /= 10 {
		result = n%10 - result
	}
	return result
}
