package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

var prog_recv = make(map[int64][]int64)
var prog_inst = make(map[int64]int)
var prog_sends = make(map[int64]int64)

func init_registers(program int64) map[string]int64 {

	registers := make(map[string]int64)
	for _, v := range([]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}) {
		registers[v] = 0
	}
	registers["p"] = program
	return registers
}

func read_file(f string) []string {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	instructions := make([]string,0)

	for s.Scan() {
		instructions = append(instructions, s.Text())
	}

	return instructions
}

func process(programid int64, registers map[string]int64, instructions []string) int {

	inst_num := prog_inst[programid]

	if inst_num >= len(instructions) {
		fmt.Println("out of range, exiting..")
		return -1
	}

		inst := strings.Split(instructions[inst_num]," ")
		instruction := inst[0]

		reg1 := inst[1]
		reg1_val := int64(0)
		if val, err := strconv.Atoi(reg1); err == nil {
			reg1_val = int64(val)
		} else {
			reg1_val = registers[reg1]
		}
		reg2_val := int64(0)
		if len(inst) > 2 {
			if val, err := strconv.Atoi(inst[2]); err == nil {
				reg2_val = int64(val)
			} else {
				reg2_val = registers[inst[2]]
			}
		}
		if instruction == "snd" {
			dest_program := int64(-1)
			if (programid == 0) {
				dest_program = 1
			} else {
				dest_program = 0
			}
			prog_recv[dest_program] = append(prog_recv[dest_program],reg1_val)
			prog_inst[programid]++
			prog_sends[programid]++
			fmt.Printf("program %d sending (send count %d)\n",programid,prog_sends[programid])
		} else if instruction == "set" {
			registers[reg1] = reg2_val
			prog_inst[programid]++
		} else if instruction == "add" {
			registers[reg1] += reg2_val
			prog_inst[programid]++
		} else if instruction == "mul" {
			registers[reg1] *= reg2_val
			prog_inst[programid]++
		} else if instruction == "mod" {
			registers[reg1] = registers[reg1] % reg2_val
			prog_inst[programid]++
		} else if instruction == "rcv" {
			if len(prog_recv[programid]) == 0 {
				return -1
			} else {
				x := prog_recv[programid][0]
				registers[reg1] = x
				prog_recv[programid] = prog_recv[programid][1:]
			}
			prog_inst[programid]++
		} else if instruction == "jgz" {
			if reg1_val > 0 {
				prog_inst[programid] = int(prog_inst[programid]) + int(reg2_val)
			} else {
				prog_inst[programid]++
			}
		} else {
			fmt.Println("wut")
		}
	return 0
}

func main() {
	inst := read_file(os.Args[1])
	p0_registers := init_registers(0)
	p1_registers := init_registers(1)
	prog_inst[0] = 0
	prog_inst[1] = 0
	prog_sends[0] = 0
	prog_sends[1] = 0
	prog_recv[0] = make([]int64,0)
	prog_recv[1] = make([]int64,0)

	for done := false; !done; {
		r0 := process(0, p0_registers, inst)
		r1 := process(1, p1_registers, inst)

		if (r0 == -1 && r1 == -1) {
			fmt.Println("we're deadlocked!")
			done = true
		}
	}
}
