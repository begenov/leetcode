package main

import (
	"fmt"
)

func main() {
	prices := []int{2, 4, 1}
	fmt.Printf("maxProfit(prices): %v\n", maxProfit(prices))
}

func maxProfit(prices []int) int {
	min_value := maxValue(prices)
	max_value := 0
	for _, v := range prices {
		if v < min_value {
			min_value = v
		} else if v-min_value > max_value {
			max_value = v - min_value
		}
	}
	return max_value
}

func maxValue(prices []int) int {
	max := 0
	for _, v := range prices {
		if v > max {
			max = v
		}
	}
	return max
}
