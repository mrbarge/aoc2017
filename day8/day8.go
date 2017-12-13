package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"fmt"
)

type Comparator string
const (
	GT Comparator = ">"
	LT Comparator = "<"
	GTE Comparator = ">="
	LTE Comparator = "<="
	EQ Comparator = "=="
	NE Comparator = "!="
)

var registers map[string]int = make(map[string]int,30)

func fill_registers(f string) {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	r_reg, _ := regexp.Compile(`(\w+).+ if (\w+) `)

	for s.Scan() {
		line := s.Text()
		matches := r_reg.FindStringSubmatch(line)
		if matches != nil {
			registers[matches[1]] = 0
			registers[matches[2]] = 0
		}
	}
}

func test_instruction(reg string, c Comparator, v int) bool {
	switch c {
	case GT:
		return registers[reg] > v
	case LT:
		return registers[reg] < v
	case GTE:
		return registers[reg] >= v
	case LTE:
		return registers[reg] <= v
	case EQ:
		return registers[reg] == v
	case NE:
		return registers[reg] != v
	}
	return false
}

func process_instructions(f string) {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	r_inst, _ := regexp.Compile(`(\w+) (inc|dec) (-?\d+) if (\w+) (.+) (-?\d+)`)

	for s.Scan() {
		line := s.Text()
		matches := r_inst.FindStringSubmatch(line)
		if matches != nil {
			op_reg := matches[1]
			op_inst := matches[2]
			op_val, _ := strconv.Atoi(matches[3])
			op_comp_reg := matches[4]
			op_comp := matches[5]
			op_comp_val, _ := strconv.Atoi(matches[6])

			if test_instruction(op_comp_reg, Comparator(op_comp), op_comp_val) {
				if op_inst == "inc" {
					registers[op_reg] += op_val
				} else if op_inst == "dec" {
					registers[op_reg] -= op_val
				} else {
					fmt.Printf("Invalid instruction: %s\n", op_inst)
				}
			}
		} else {
			fmt.Println("Did not match")
		}
	}
}

func find_largest_val() int {
	largest := 0
	for _, v := range(registers) {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	datafile := os.Args[1]
	fill_registers(datafile)
	process_instructions(datafile)
	largest := find_largest_val()
	fmt.Printf("Largest value: %d", largest)
}

