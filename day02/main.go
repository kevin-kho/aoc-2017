package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type MinMax struct {
	Min int
	Max int
}

func (m MinMax) GetDifference() int {
	return m.Max - m.Min
}

func GetRows(data []byte) ([][]int, error) {
	var rows [][]int
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		var row []int
		for val := range strings.SplitSeq(string(entry), "\t") {
			if val == "" {
				continue
			}
			v, err := strconv.Atoi(val)
			if err != nil {
				return rows, err
			}
			row = append(row, v)
		}
		rows = append(rows, row)

	}

	return rows, nil
}

func GetMinMax(row []int) MinMax {
	minimum := math.MaxInt
	maximum := math.MinInt
	for _, v := range row {
		minimum = min(minimum, v)
		maximum = max(maximum, v)
	}

	return MinMax{
		Min: minimum,
		Max: maximum,
	}
}

func GetDivisble(row []int) (int, int, error) {
	for i, a := range row {
		for j := i + 1; j < len(row); j++ {
			b := row[j]
			if a == b {
				return a, b, nil
			}
			if a > b && a%b == 0 {
				return a, b, nil
			}

			if a < b && b%a == 0 {
				return b, a, nil
			}
		}
	}

	return 0, 0, errors.New("no divisible pair")

}

func SolvePartOne(rows [][]int) int {

	var minMaxs []MinMax
	for _, row := range rows {
		minMaxs = append(minMaxs, GetMinMax(row))
	}
	var res int
	for _, m := range minMaxs {
		res += m.GetDifference()
	}
	return res

}

func SolvePartTwo(rows [][]int) (int, error) {
	var res int
	for _, row := range rows {
		a, b, err := GetDivisble(row)
		if err != nil {
			return res, err
		}
		res += a / b
	}
	return res, nil
}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	rows, err := GetRows(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(rows)
	fmt.Println(res)

	res2, err := SolvePartTwo(rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res2)

}
