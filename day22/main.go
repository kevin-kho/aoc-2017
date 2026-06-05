package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Virus struct {
	Pos
	Dx int
	Dy int
}

type Pos struct {
	X int
	Y int
}

func (v *Virus) TurnLeft() {

	curr := []int{v.Dx, v.Dy}

	if slices.Equal(curr, []int{0, 1}) {
		v.Dx = -1
		v.Dy = 0
	} else if slices.Equal(curr, []int{-1, 0}) {
		v.Dx = 0
		v.Dy = -1
	} else if slices.Equal(curr, []int{0, -1}) {
		v.Dx = 1
		v.Dy = 0
	} else if slices.Equal(curr, []int{1, 0}) {
		v.Dx = 0
		v.Dy = 1
	}

}

func (v *Virus) TurnRight() {
	curr := []int{v.Dx, v.Dy}

	if slices.Equal(curr, []int{0, 1}) {
		v.Dx = 1
		v.Dy = 0
	} else if slices.Equal(curr, []int{-1, 0}) {
		v.Dx = 0
		v.Dy = 1
	} else if slices.Equal(curr, []int{0, -1}) {
		v.Dx = -1
		v.Dy = 0
	} else if slices.Equal(curr, []int{1, 0}) {
		v.Dx = 0
		v.Dy = -1
	}
}

func (v *Virus) Move() {
	v.X += v.Dx
	v.Y += v.Dy
}

type Grid [][]byte

func (g Grid) InitializeVirus() Virus {

	y := len(g) / 2
	x := len(g[0]) / 2

	return Virus{
		Pos: Pos{
			X: x,
			Y: y,
		},
		Dx: 0,
		Dy: 1,
	}

}

func (g Grid) GetInfected() map[Pos]bool {
	res := make(map[Pos]bool)

	for y, row := range g {
		for x, node := range row {
			if node == '#' {
				res[Pos{X: x, Y: y}] = true
			}

		}
	}

	return res

}

func CreateGrid(data []byte) Grid {
	return bytes.Split(data, []byte{'\n'})
}

func SolvePartOne(virus Virus, infected map[Pos]bool) int {
	var count int

	for range 10000 {

		// Infected Node
		if infected[virus.Pos] {
			infected[virus.Pos] = false
			virus.TurnRight()

		} else {
			// Non infected node
			infected[virus.Pos] = true
			virus.TurnLeft()

			count++

			// fmt.Printf("infected %v\n", virus.Pos)
		}

		// Move
		virus.Move()

	}

	for pos, b := range infected {
		if b {
			fmt.Println(pos)
		}
	}

	fmt.Println(virus)

	return count

}

func main() {

	data, err := common.ReadInput("./inputExample.txt")
	// data, err := common.ReadInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	grid := CreateGrid(data)

	virus := grid.InitializeVirus()
	infected := grid.GetInfected()

	fmt.Println(virus)
	fmt.Println(infected)

	res := SolvePartOne(virus, infected)
	fmt.Println(res)

}
