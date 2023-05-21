package main

import "fmt"

func main() {
	arr := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr))
}

func validMountainArray(arr []int) bool {
	var a []int
	var b []int

	for i := range arr {
		if arr[i] < arr[i+1] {
			fmt.Println(i)
			a = append(a, arr[i])
		} else if arr[i] == arr[i+1] {
			return false
		} else {
			break
		}
	}
	for i := len(a); i < len(arr); i++ {
		if arr[i] < arr[i+1] {
			b = append(b, arr[i])
		} else if arr[i] == arr[i+1] {
			return false
		} else {
			break
		}
	}

	fmt.Println(a, "a - b", b)
	return true
}
