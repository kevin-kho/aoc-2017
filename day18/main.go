package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Instruction struct {
	Command  string // snd, set, add, mul, mod, rcv jgz
	Register string
	Value    int
	ValueReg string
}

func GetInstructions(data []byte) ([]Instruction, error) {
	var res []Instruction

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStr := string(entry)

		entryStrArr := strings.Split(entryStr, " ")
		cmd := entryStrArr[0]
		var register string
		var val int
		var valReg string
		var err error

		switch len(entryStrArr) {
		case 2:
			val, err = strconv.Atoi(entryStrArr[1])
			if err != nil {
				valReg = entryStrArr[1]
			}
		case 3:
			register = entryStrArr[1]
			val, err = strconv.Atoi(entryStrArr[2])
			if err != nil {
				valReg = entryStrArr[2]
			}
		}

		ins := Instruction{
			Command:  cmd,
			Register: register,
			Value:    val,
			ValueReg: valReg,
		}

		res = append(res, ins)
	}

	return res, nil

}

func SolvePartOne(ins []Instruction) int {
	reg := make(map[string]int)
	i := 0

	var played []int

	for 0 <= i && i < len(ins) {
		in := ins[i]

		var value int
		if in.ValueReg != "" {
			value = reg[in.ValueReg]
		} else {
			value = in.Value
		}

		switch in.Command {
		case "snd":
			played = append(played, value)
		case "set":
			reg[in.Register] = value
		case "add":
			reg[in.Register] += value
		case "mul":
			reg[in.Register] *= value
		case "mod":
			reg[in.Register] %= value
		case "rcv":
			if value != 0 {
				// Handle case that there's empty slice
				return played[len(played)-1]
			}
		case "jgz":

			// Hacky; need to redo CreateInstructions
			v, err := strconv.Atoi(in.Register)
			if err != nil {
				v = reg[in.Register]
			}

			if v > 0 {
				i += value
				continue
			}

		}

		i++

	}

	return 0

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	ins, err := GetInstructions(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(ins)
	fmt.Println(res)

}
