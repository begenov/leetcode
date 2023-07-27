package main

import "fmt"

func main() {
	num := 199

	fmt.Printf("addDigits(num): %v\n", addDigits(num))
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	var res int
	for num != 0 {
		add := num % 10
		num /= 10
		res += add
	}

	return addDigits(res)

}
