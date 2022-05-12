package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "abcabcbb"
	fmt.Println(lengthOfLongestSubString(input))

}

func lengthOfLongestSubString(input string) int {
	var length int
	maxLength := 0
	prevPosition := -1
	for i, v := range input {
		c := string(v)
		position := strings.LastIndex(input[:i], c)
		if position > prevPosition {
			prevPosition = position
		}
		length = i - prevPosition
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
