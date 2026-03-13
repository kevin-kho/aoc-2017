package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func ConvertIntToRunes(i int) []rune {
	var res []rune
	for _, c := range strconv.Itoa(i) {
		res = append(res, c)
	}

	return res
}

func RunRounds(input []int, rounds int) []int {
	// Build list
	var list []int
	for i := range 256 {
		list = append(list, i)
	}

	var start int
	var skip int

	for range rounds {

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

			start = (start + in + skip) % 256

			skip++

		}
	}

	return list

}

func SolvePartOne(input []int) int {

	list := RunRounds(input, 1)

	return list[0] * list[1]

}

func SolvePartTwo(input []int) string {
	l1 := RunRounds(input, 64)

	var res []int
	for i := 0; i < len(l1); i += 16 {
		var num int
		for j := i; j < i+16; j++ {
			num ^= l1[j]
		}
		res = append(res, num)

	}

	var sb strings.Builder

	for _, n := range res {
		hex := fmt.Sprintf("%02x", n)
		sb.WriteString(hex)
	}

	return sb.String()

}

func main() {

	input := []int{130, 126, 1, 11, 140, 2, 255, 207, 18, 254, 246, 164, 29, 104, 0, 224}
	res := SolvePartOne(input)
	fmt.Println(res)

	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)
	data = append(data, 17, 31, 73, 47, 23)

	var input2 []int
	for _, b := range data {
		input2 = append(input2, int(b))
	}

	res2 := SolvePartTwo(input2)
	fmt.Println(res2)

}
