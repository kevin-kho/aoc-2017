package main

import (
	"fmt"
	"slices"
)

func SolvePartOne() int {
	arr := []int{0}
	val := 1
	i := 0
	for range 2017 {
		i += 376
		i %= len(arr)

		arr = slices.Insert(arr, i+1, val)
		i++
		val++

	}

	if i == 2017 {
		return arr[0]
	}

	return arr[i+1]
}

func main() {
	res := SolvePartOne()
	fmt.Println(res)

}
