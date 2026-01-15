package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func CreateIntArr(data []byte) ([]int, error) {
	var arr []int
	for v := range bytes.SplitSeq(data, []byte{'\n'}) {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			return arr, err
		}
		arr = append(arr, i)
	}

	return arr, nil
}

func SolvePartOne(instructions []int) int {
	var steps int
	var i int
	for i < len(instructions) {
		newI := i + instructions[i]
		instructions[i]++

		i = newI
		steps++

	}

	return steps
}

func SolvePartTwo(instructions []int) int {
	var steps int
	var i int
	for i < len(instructions) {
		newI := i + instructions[i]
		if instructions[i] >= 3 {
			instructions[i]--
		} else {
			instructions[i]++
		}

		i = newI
		steps++

	}

	return steps
}

func main() {

	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArrPartOne, err := CreateIntArr(data)
	if err != nil {
		log.Fatal(err)
	}
	intArrPartTwo := slices.Clone(intArrPartOne)

	res := SolvePartOne(intArrPartOne)
	fmt.Println(res)

	res2 := SolvePartTwo(intArrPartTwo)
	fmt.Println(res2)

}
