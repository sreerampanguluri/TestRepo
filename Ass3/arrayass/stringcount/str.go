package main

import (
	"fmt"
	"strings"
)

func repetition(st string) map[string]int {

	s := strings.Fields(st)
	wordcount := make(map[string]int)
	for _, word := range s {
		_, matched := wordcount[word]
		if matched {
			wordcount[word] += 1
		} else {
			wordcount[word] = 1
		}
	}
	return wordcount
}

func main() {
	s := "I scream, you scream, we all scream for ice cream , "
	for index, element := range repetition(s) {
		fmt.Println(index, "=", element)
	}
}
