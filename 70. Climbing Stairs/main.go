package main

import "fmt"

func main() {
	n := 45

	fmt.Printf("climbStairs(n): %v\n", climbStairs(n))
}

func climbStairs(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
