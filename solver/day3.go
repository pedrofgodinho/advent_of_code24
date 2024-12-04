package solver

import (
	"fmt"
	"regexp"
	"strconv"
)

// enum of instructions
const (
	mul = iota
	do
	dont
)

type Instruction struct {
	instruction int
	left        int
	right       int
}

type Day3 struct {
	instructions []Instruction
}

var regex = regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)

func (d *Day3) Parse(input string) (bool, error) {
	match := regex.FindAllStringSubmatch(input, -1)
	for _, m := range match {
		switch m[0] {
		case "do()":
			d.instructions = append(d.instructions, Instruction{do, 0, 0})
		case "don't()":
			d.instructions = append(d.instructions, Instruction{dont, 0, 0})
		default:
			left, _ := strconv.Atoi(m[2])
			right, _ := strconv.Atoi(m[3])
			d.instructions = append(d.instructions, Instruction{mul, left, right})
		}
	}
	return true, nil
}

func (d *Day3) Part1() (string, error) {
	sum := 0
	for _, m := range d.instructions {
		if m.instruction == mul {
			sum += m.left * m.right
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func (d *Day3) Part2() (string, error) {
	sum := 0
	enabled := true
	for _, m := range d.instructions {
		switch m.instruction {
		case mul:
			if enabled {
				sum += m.left * m.right
			}
		case do:
			enabled = true
		case dont:
			enabled = false
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
