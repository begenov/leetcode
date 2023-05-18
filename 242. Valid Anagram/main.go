package main

import "log"

func main() {
	s := "anagram"
	t := "nagaram"
	log.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	c := make(map[rune]int)
	j := make(map[rune]int)
	for _, v := range s {
		c[v]++
	}
	for _, v := range t {
		j[v]++
	}

	for key, value := range c {
		if jValue, ok := j[key]; !ok || value != jValue {
			return false
		}
	}

	return true
}
