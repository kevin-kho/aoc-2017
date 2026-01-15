package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SumNextMatchingDigits(data []byte) int {
	var res int
	for i := 1; i < len(data); i++ {
		if data[i-1] == data[i] {
			res += int(data[i] - '0')
		}
	}

	if data[0] == data[len(data)-1] {
		res += int(data[0] - '0')
	}

	return res
}

func SumNextHalfwayDigit(data []byte) int {
	l := len(data)
	n := l / 2
	var res int
	for i := range data {
		j := (i + n) % len(data)
		if data[i] == data[j] {
			res += int(data[i] - '0')
		}
	}
	return res
}

func SolvePartOne(data []byte) int {
	return SumNextMatchingDigits(data)
}

func SolvePartTwo(data []byte) int {
	return SumNextHalfwayDigit(data)
}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	res := SolvePartOne(data)
	fmt.Println(res)

	res2 := SolvePartTwo(data)
	fmt.Println(res2)

}
