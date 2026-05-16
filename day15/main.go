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

func main() {

	res := SolvePartOne(679, 771)

	fmt.Println(res)

}
