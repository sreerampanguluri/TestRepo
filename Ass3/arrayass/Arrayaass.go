package main

import "fmt"

func main() {
	var givenarray = []int{2, 3, 4, 5, 6, 6, 7, 7, 8, 8, 2, 4}
	var x, y int
	var newarray []int

	fmt.Println("Find Duplicate number in a array")

	fmt.Printf("The given list of number as %v\n", givenarray)

	for x = 0; x < len(givenarray); x++ {
		for y = x + 1; y < len(givenarray); y++ {
			if givenarray[x] == givenarray[y] {
				newarray = append(newarray, givenarray[x])
			}
		}
	}
}
