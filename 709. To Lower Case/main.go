package main

import "log"

func main() {
	s := "Hello World"

	log.Println(toLowerCase(s))
}

func toLowerCase(s string) string {
	var res string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			res += string(v + 32)
		} else {
			res += string(v)
		}
	}
	return res
}
