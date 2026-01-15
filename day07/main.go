package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Disc struct {
	Name     string
	Weight   int
	Children []Disc
}

func CreateAdjList(data []byte) map[string][]string {
	adj := make(map[string][]string)
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		if strings.Contains(string(entry), "->") {
			strArr := strings.Split(string(entry), " -> ")
			key := strings.Split(strArr[0], " ")[0]
			values := strings.Split(strArr[1], ", ")

			adj[key] = values

		}
	}

	return adj
}

func CreateWeightMap(data []byte) (map[string]int, error) {
	wt := make(map[string]int)

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStr := string(entry)
		key := strings.Split(entryStr, " ")[0]
		i := strings.Index(entryStr, "(")
		j := strings.Index(entryStr, ")")

		w, err := strconv.Atoi(entryStr[i+1 : j])
		if err != nil {
			return wt, err
		}
		wt[key] = w

	}

	return wt, nil
}

func FindRoot(adj map[string][]string) string {
	isChild := make(map[string]bool)
	for _, children := range adj {
		for _, c := range children {
			isChild[c] = true
		}
	}

	for k := range adj {
		if !isChild[k] {
			return k
		}
	}

	return ""

}

func ConstructNaryTree(rootName string, adj map[string][]string, wt map[string]int) Disc {

	var construct func(discName string) Disc
	construct = func(discName string) Disc {

		d := Disc{
			Name:     discName,
			Weight:   wt[discName],
			Children: []Disc{},
		}
		for _, c := range adj[discName] {
			d.Children = append(d.Children, construct(c))
		}
		return d

	}

	return construct(rootName)

}

func CalculateTreeWeight(disc Disc) int {
	var res int
	queue := []Disc{disc}
	for len(queue) > 0 {
		for range len(queue) {
			d := queue[0]
			queue = queue[1:]

			res += d.Weight
			for _, c := range d.Children {
				queue = append(queue, c)
			}
		}
	}

	return res
}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	adj := CreateAdjList(data)
	wt, err := CreateWeightMap(data)
	if err != nil {
		log.Fatal(err)
	}

	rootName := FindRoot(adj)
	fmt.Println(rootName)

	tree := ConstructNaryTree(rootName, adj, wt)
	fmt.Println(tree)

}
