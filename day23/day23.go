package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func init_registers() map[string]int {

	registers := make(map[string]int)
	for _, v := range([]string{"a","b","c","d","e","f","g","h"}) {
		registers[v] = 0
	}
	return registers
}

func read_file(f string) {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	registers := init_registers()
	instructions := make([]string,0)

	for s.Scan() {
		instructions = append(instructions, s.Text())
	}

	num_mul := 0
	for i := 0; i < len(instructions); i++ {
		fmt.Println("* ",i,registers)
		inst := strings.Split(instructions[i]," ")
		instruction := inst[0]

		reg1 := inst[1]
		reg1_val := 0
		if val, err := strconv.Atoi(reg1); err == nil {
			reg1_val = val
		} else {
			reg1_val = registers[reg1]
		}

		reg2_val := 0
		if len(inst) > 2 {
			if val, err := strconv.Atoi(inst[2]); err == nil {
				reg2_val = val
			} else {
				reg2_val = registers[inst[2]]
			}
		}
		if instruction == "set" {
			registers[reg1] = reg2_val
		} else if instruction == "sub" {
			registers[reg1] -= reg2_val
		} else if instruction == "mul" {
			registers[reg1] *= reg2_val
			num_mul++
		} else if instruction == "jnz" {
			if reg1_val != 0 {
				i = i + reg2_val - 1
			}
		}
	}
	fmt.Println(num_mul)

}

func main() {
	read_file(os.Args[1])
}
