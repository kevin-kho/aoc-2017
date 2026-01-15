package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

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

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	adj := CreateAdjList(data)

	res := FindRoot(adj)
	fmt.Println(res)
}
