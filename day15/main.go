package main

import (
	"fmt"
	"strconv"
)

func SolvePartOne(a int, b int) int {
	var score int

	aFactor := 16807
	bFactor := 48271

	for range 40_000_000 {
		a *= aFactor
		a %= 2147483647

		b *= bFactor
		b %= 2147483647

		aBin := strconv.FormatInt(int64(a), 2)
		aBin = fmt.Sprintf("%032s", aBin)

		bBin := strconv.FormatInt(int64(b), 2)
		bBin = fmt.Sprintf("%032s", bBin)

		if aBin[16:] == bBin[16:] {
			score++
		}
	}

	return score

}

func SolvePartTwo(a int, b int) int {
	var score int

	aFactor := 16807
	bFactor := 48271

	var aValues []int
	for len(aValues) < 5_000_000 {
		a *= aFactor
		a %= 2147483647
		if a%4 == 0 {
			aValues = append(aValues, a)
		}
	}

	var bValues []int
	for len(bValues) < 5_000_000 {
		b *= bFactor
		b %= 2147483647
		if b%8 == 0 {
			bValues = append(bValues, b)
		}
	}

	for i := range 5_000_000 {
		a = aValues[i]
		b = bValues[i]

		aBin := strconv.FormatInt(int64(a), 2)
		aBin = fmt.Sprintf("%032s", aBin)

		bBin := strconv.FormatInt(int64(b), 2)
		bBin = fmt.Sprintf("%032s", bBin)

		if aBin[16:] == bBin[16:] {
			score++
		}

	}

	return score

}

func main() {

	var res int
	// res = SolvePartOne(679, 771)
	res = SolvePartTwo(679, 771)

	fmt.Println(res)

}
