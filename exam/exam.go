package main

import (
	"fmt"
)

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	Slice := []int{143, 02, 66, 111, 0, 66, 69, 9, 18, 18, 81, 81, 143, 143}
	fmt.Println(Slice)
	newSlice := unique(Slice)
	fmt.Println(newSlice)
}
