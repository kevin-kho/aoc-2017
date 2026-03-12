package main

import (
	"fmt"
	"log"
	"os"
)

func SolvePartOne(data []byte) int {
	// while loop through t
	i := 0

	var braces []byte
	var collecting bool

	var score int
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

	return score
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(data)
	fmt.Println(res)

}
