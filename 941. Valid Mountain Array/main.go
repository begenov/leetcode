package main

import "fmt"

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr))
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	i := 0

	for ; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println(i)
			i++
			break
		}
		if arr[i] == arr[i+1] {
			return false
		}
	}

	for ; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			return false
		}
	}

	return true

}
