package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	result := 0
	for idx, number := range numbers{
		fmt.Println(idx, number)
		result += number
	}
	return result
}

func main() {
	total := superAdd(4, 3, 1, 2, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
}