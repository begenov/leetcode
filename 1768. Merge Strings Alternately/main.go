package main

import (
	"log"
)

func main() {
	word1 := "ab"
	word2 := "pqrs"
	log.Println(mergeAlternately(word1, word2))
}
func mergeAlternately(word1 string, word2 string) string {
	var i, j int
	var res string
	for i < len(word1) && j < len(word2) {
		res += string(word1[i])
		res += string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		res += string(word1[i])
		i++
	}

	for j < len(word2) {
		res += string(word2[j])
		j++
	}

	return res

}
