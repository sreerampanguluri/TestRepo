package main

import "fmt"

func duplicateInArray(arr []int) int {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return 0
}

func main() {
	fmt.Println(duplicateInArray([]int{1, 2, 3, 3, 4}))
	fmt.Println(duplicateInArray([]int{2, 5, 6, 3, 5}))
	fmt.Println(duplicateInArray([]int{1, 2, 3, 4, 5}))
}
