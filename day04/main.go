package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func CreatePassPhrases(data []byte) [][]string {
	var res [][]string
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		res = append(res, strings.Split(string(entry), " "))
	}

	return res
}

func IsValidPartOne(passPhrase []string) bool {
	seen := make(map[string]bool)
	for _, p := range passPhrase {
		if seen[p] {
			return false
		}
		seen[p] = true
	}

	return true
}

func IsValidPartTwo(passPhrase []string) bool {
	seen := make(map[string]bool)
	for _, p := range passPhrase {
		b := []byte(p)
		slices.Sort(b)
		if seen[string(b)] {
			return false
		}
		seen[string(b)] = true

	}

	return true

}

func SolvePartOne(passPhrases [][]string) int {
	var count int
	for _, p := range passPhrases {
		if IsValidPartOne(p) {
			count++
		}
	}
	return count
}

func SolvePartTwo(passPhrases [][]string) int {
	var count int
	for _, p := range passPhrases {
		if IsValidPartTwo(p) {
			count++
		}
	}
	return count
}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	passPhrases := CreatePassPhrases(data)

	res := SolvePartOne(passPhrases)
	fmt.Println(res)

	res2 := SolvePartTwo(passPhrases)
	fmt.Println(res2)

}
