package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"
	"slices"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Layer struct {
	Depth int
	Range int
	Curr  int
	Dir   int
}

func (l *Layer) Move(x int) {
	for range x {
		l.Curr += l.Dir

		if l.Curr == 0 || l.Curr == l.Range-1 {
			l.Dir *= -1
		}

	}

}

func (l Layer) GetScore() int {
	return l.Depth * l.Range
}

func CreateLayers(data []byte) ([]Layer, error) {

	var res []Layer
	mp := make(map[int]int)

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		depth := bytes.Split(entry, []byte{':'})[0]
		rng := bytes.Split(entry, []byte{':'})[1]

		rng = bytes.TrimSpace(rng)

		d, err := strconv.Atoi(string(depth))
		if err != nil {
			return res, err
		}

		r, err := strconv.Atoi(string(rng))
		if err != nil {
			return res, err
		}

		mp[d] = r

	}

	keys := slices.Sorted(maps.Keys(mp))
	st := keys[0]
	ed := keys[len(keys)-1]

	for i := st; i <= ed; i++ {
		l := Layer{
			Depth: i,
			Range: mp[i],
			Dir:   1,
		}

		res = append(res, l)
	}

	return res, nil

}

func SolvePartOne(layers []Layer) int {

	var score int

	for i, l := range layers {
		l.Move(i)
		if l.Curr == 0 {
			score += l.GetScore()
		}

	}

	return score

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	layers, err := CreateLayers(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(layers)
	fmt.Println(res)

}
