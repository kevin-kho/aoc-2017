package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SumMatchingDigits(data []byte) int {
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

func SolvePartOne(data []byte) int {
	return SumMatchingDigits(data)
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

}
