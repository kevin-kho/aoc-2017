package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

func GetDirPos(dir string) Pos {
	var d Pos
	switch dir {
	case "n":
		d.X = 0
		d.Y = 1
	case "ne":
		d.X = 1
		d.Y = 1

	case "se":
		d.X = 1
		d.Y = -1

	case "s":
		d.X = 0
		d.Y = -1

	case "sw":
		d.X = -1
		d.Y = -1

	case "nw":
		d.X = -1
		d.Y = 1

	}

	return d
}

func CreatePos(data []byte) []Pos {

	var res []Pos

	for entry := range bytes.SplitSeq(data, []byte{','}) {
		res = append(res, GetDirPos(string(entry)))
	}

	return res
}

func SolvePartOne(dirs []Pos) int {
	var curr Pos

	for _, d := range dirs {
		curr.X += d.X
		curr.Y += d.Y
	}

	fmt.Println(curr)

	x := common.IntAbs(curr.X)
	y := common.IntAbs(curr.Y)

	fmt.Println(x, y)

	var steps int
	sub := min(x, y)
	steps += sub
	x -= sub
	y -= sub

	if x != 0 {
		steps += x
	}

	if y != 0 {
		steps += y
	}

	return steps

}

func main() {

	data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)
	dirs := CreatePos(data)

	res := SolvePartOne(dirs)
	fmt.Println(res)

}
