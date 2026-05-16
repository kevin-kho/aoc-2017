package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Cast struct {
	Programs []string
	IndexOf  map[string]int
}

func (c *Cast) Move(m Move) {
	switch m.Cmd {
	case "Spin":
		for p, i := range c.IndexOf {
			i += m.Value
			i %= len(c.Programs)
			c.IndexOf[p] = i

			c.Programs[i] = p
		}

	case "Exchange":
		i := m.Indexes[0]
		j := m.Indexes[1]

		c.Programs[i], c.Programs[j] = c.Programs[j], c.Programs[i]
		for i, p := range c.Programs {
			c.IndexOf[p] = i
		}
	case "Partner":

		i := c.IndexOf[m.Partners[0]]
		j := c.IndexOf[m.Partners[1]]

		c.Programs[i], c.Programs[j] = c.Programs[j], c.Programs[i]
		for i, p := range c.Programs {
			c.IndexOf[p] = i
		}
	}
}

func (c Cast) GetOrder() string {
	return strings.Join(c.Programs, "")
}

func CreateCast(cast string) Cast {
	var programs []string
	programIdx := make(map[string]int)
	for i, char := range cast {
		programs = append(programs, string(char))
		programIdx[string(char)] = i
	}
	return Cast{
		Programs: programs,
		IndexOf:  programIdx,
	}
}

type Move struct {
	Cmd      string
	Indexes  []int
	Partners []string
	Value    int
}

func CreateMoves(data []byte) ([]Move, error) {
	var moves []Move
	for entry := range bytes.SplitSeq(data, []byte{','}) {

		var cmd string
		var indexes []int
		var partners []string
		var value int
		var err error

		prefix := entry[0]
		suffix := entry[1:]

		switch prefix {
		case 's':
			cmd = "Spin"
			value, err = strconv.Atoi(string(suffix))
			if err != nil {
				return moves, err
			}

		case 'x':
			cmd = "Exchange"
			for idxStr := range strings.SplitSeq(string(suffix), "/") {
				idx, err := strconv.Atoi(idxStr)
				if err != nil {
					return moves, err
				}
				indexes = append(indexes, idx)

			}

		case 'p':
			cmd = "Partner"
			partners = strings.Split(string(suffix), "/")
		}

		move := Move{
			Cmd:      cmd,
			Indexes:  indexes,
			Partners: partners,
			Value:    value,
		}
		moves = append(moves, move)

	}

	return moves, nil

}

func SolvePartOne(cast Cast, moves []Move) string {
	for _, m := range moves {
		cast.Move(m)
	}

	return cast.GetOrder()

}

func SolvePartTwo(cast Cast, moves []Move) string {
	// Too slow
	seen := make(map[string][]string)
	for range 1_000_000_000 {
		start := cast.GetOrder()
		if end, ok := seen[start]; ok {
			cast.Programs = slices.Clone(end)
			for i, p := range cast.Programs {
				cast.IndexOf[p] = i
			}
			continue
		}

		for _, m := range moves {
			cast.Move(m)
		}

		seen[start] = slices.Clone(cast.Programs)

	}

	return cast.GetOrder()
}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	moves, err := CreateMoves(data)
	if err != nil {
		log.Fatal(err)
	}

	castPartOne := CreateCast("abcdefghijklmnop")
	castPartTwo := CreateCast("abcdefghijklmnop")

	res := SolvePartOne(castPartOne, moves)
	fmt.Println(res)

	res2 := SolvePartTwo(castPartTwo, moves)
	fmt.Println(res2)

}
