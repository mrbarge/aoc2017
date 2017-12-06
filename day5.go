package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func readdata(f string) []int {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	var r []int = make([]int,0)

	for s.Scan() {
		inst, _ := strconv.Atoi(s.Text())
		r = append(r,inst)
	}
	return r
}

func run_instructions(instructions []int) int {
	escaped := false

	i := 0
	steps := 0
	num_instructions := len(instructions)
	for !escaped {
		inst := instructions[i]
		instructions[i]++
		i = i + inst
		steps++
		if (i < 0 || i >= num_instructions) {
			escaped = true
		}
	}
	return steps
}

func main() {
	datafile := os.Args[1]
	data := readdata(datafile)
	steps := run_instructions(data)
	fmt.Println(steps)
}
