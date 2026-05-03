package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func ConvertStringToInt(s string) (int, error) {

	s = strings.TrimSpace(s)
	sInt, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return sInt, nil

}

func CreateAdjList(data []byte) (map[int][]int, error) {
	res := make(map[int][]int)

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStr := string(entry)
		entryStrArr := strings.Split(entryStr, "<->")

		src := entryStrArr[0]
		dsts := entryStrArr[1]
		dstsStrArr := strings.Split(dsts, ",")

		srcInt, err := ConvertStringToInt(src)
		if err != nil {
			log.Fatal(err)
		}

		var dstInt []int
		for _, d := range dstsStrArr {
			dInt, err := ConvertStringToInt(d)
			if err != nil {
				log.Fatal(err)
			}
			dstInt = append(dstInt, dInt)

		}

		res[srcInt] = append(res[srcInt], dstInt...)
		for _, d := range dstInt {
			res[d] = append(res[d], srcInt)
		}

	}

	// Keep unique dst
	for k, v := range res {
		unique := make(map[int]bool)
		for _, n := range v {
			unique[n] = true
		}

		res[k] = slices.Sorted(maps.Keys(unique))
	}

	return res, nil

}

func SolvePartOne(adjList map[int][]int) int {

	visited := make(map[int]bool)

	// Plan to DFS from 0 outwards
	// Keep visited in a set
	var dfs func(node int)
	dfs = func(node int) {

		// case: visited
		if visited[node] {
			return
		}

		// case visit and go next neighbors
		visited[node] = true
		for _, neighbor := range adjList[node] {
			dfs(neighbor)
		}

	}

	dfs(0)

	return len(visited)

}

func SolvePartTwo(adjList map[int][]int) int {

	visited := make(map[int]bool)

	// Plan to DFS from 0 outwards
	// Keep visited in a set
	var dfs func(node int)
	dfs = func(node int) {

		// case: visited
		if visited[node] {
			return
		}

		// case visit and go next neighbors
		visited[node] = true
		for _, neighbor := range adjList[node] {
			dfs(neighbor)
		}

	}

	var groups int
	for n, _ := range adjList {

		before := len(visited)
		dfs(n)
		after := len(visited)
		if after > before {
			groups++
		}

	}

	return groups

}

func main() {
	// data, err := common.ReadInput("./inputExample.txt")
	data, err := common.ReadInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	adjList, err := CreateAdjList(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(adjList)
	fmt.Println(res)

	res2 := SolvePartTwo(adjList)
	fmt.Println(res2)

}
