package main

import "fmt"

func main() {
	s := "a"
	fmt.Printf("lengthOfLastWord(s): %v\n", lengthOfLastWord(s))
}
func lengthOfLastWord(s string) int {
	var word string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && len(word) == 0 {
			continue
		}

		if s[i] != ' ' {
			word += string(s[i])
		} else {
			break
		}
	}
	return len(word)
}
