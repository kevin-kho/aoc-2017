package main

import (
	"fmt"
	"log"
	"os"
)

func Solve(data []byte) (int, int) {
	// while loop through t
	i := 0

	var braces []byte
	var collecting bool
	var garbage []byte

	var score int
	var totalGarbage int
	for i < len(data) {
		char := data[i]

		if char == '!' {
			i += 2
			continue
		}

		// garbage collecting
		if collecting {

			// close bag, stop collecting
			if char == '>' {
				collecting = false
				totalGarbage += len(garbage)
				garbage = []byte{}
			}
			if collecting {
				garbage = append(garbage, char)
			}
			i++
			continue
		}

		// not garbage collecting
		if char == '{' {
			braces = append(braces, '{')
		}

		if char == '}' {
			score += len(braces)
			braces = braces[:len(braces)-1]
		}

		if char == '<' {
			collecting = true
		}

		i++

	}

	return score, totalGarbage
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	resOne, resTwo := Solve(data)
	fmt.Println(resOne, resTwo)

}
