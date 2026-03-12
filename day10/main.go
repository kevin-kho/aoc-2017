package main

import (
	"fmt"
	"slices"
)

func SolvePartOne(input []int) int {

	// Build list
	var list []int
	for i := range 256 {
		list = append(list, i)
	}

	var start int
	var skip int

	for _, in := range input {
		var arr []int
		for i := start; i < start+in; i++ {
			index := i % 256
			arr = append(arr, list[index])
		}
		slices.Reverse(arr)
		j := 0
		for i := start; i < start+in; i++ {
			index := i % 256
			list[index] = arr[j]
			j++
		}

		start = (start + in) % 256
		start += skip

		skip++

	}

	return list[0] * list[1]

}

func main() {

	input := []int{130, 126, 1, 11, 140, 2, 255, 207, 18, 254, 246, 164, 29, 104, 0, 224}
	res := SolvePartOne(input)
	fmt.Println(res)

}
