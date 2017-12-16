package main

import (
	"flag"
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
	"regexp"
)

func spin(positions []string, spins int) {
	// grab the tail
	tail := positions[(len(positions)-spins):]
	// grab the head
	head := positions[0:(len(positions)-spins)]
	// swap em
	new_pos := append(tail,head...)
	// fill in the original positions
	for i, _ := range(new_pos) {
		positions[i] = new_pos[i]
	}
}

func exchange(positions []string, pos1 int, pos2 int) {
	if (pos1 > len(positions) || pos2 > len(positions) ||
		pos1 < 0 || pos2 < 0) {
			fmt.Printf("invalid instruction to swap %d and %d\n",pos1,pos2)
			return
	}
	tmppos := positions[pos1]
	positions[pos1] = positions[pos2]
	positions[pos2] = tmppos
}

func find_pos(positions []string, prog string) int {
	for i, v := range(positions) {
		if v == prog {
			return i
		}
	}
	return -1
}

func partner(positions []string, prog1 string, prog2 string) {
	pos1 := find_pos(positions, prog1)
	pos2 := find_pos(positions, prog2)
	if (pos1 < 0 || pos2 < 0) {
		fmt.Println("Unable to find program pair %s/%s\n",prog1,prog2)
		return
	}
	exchange(positions,pos1,pos2)
}

func perform_dance(positions []string, instructions []string) {

	rspin, _ := regexp.Compile(`s(\d+)`)
	rexch, _ := regexp.Compile(`x(\d+)/(\d+)`)
	rpart, _ := regexp.Compile(`p(.+)/(.+)`)

	for _, instruction := range(instructions) {

		if instruction[0] == 's' {
			lm := rspin.FindStringSubmatch(instruction)
			if lm != nil {
				spins, _ := strconv.Atoi(string(lm[1]))
				spin(positions, spins)
			}
		} else if instruction[0] == 'x' {
			lm := rexch.FindStringSubmatch(instruction)
			if lm != nil {
				pos1, _ := strconv.Atoi(lm[1])
				pos2, _ := strconv.Atoi(lm[2])
				exchange(positions, pos1, pos2)
			}
		} else if instruction[0] == 'p' {
				lm := rpart.FindStringSubmatch(instruction)
				if lm != nil {
					prog1 := lm[1]
					prog2 := lm[2]
					partner(positions, prog1, prog2)
				}
		}
	}
}

func get_instructions(filepath string) []string {
	file, _ := os.Open(filepath)
	s := bufio.NewScanner(file)
	s.Scan()
	dance := s.Text()

	return strings.Split(dance, ",")
}

func setup_initial_positions() []string {
	return []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p"}
}

func main() {

	infilePtr := flag.String("input", "day16.txt", "input file")

	instructions := get_instructions(*infilePtr)
	positions := setup_initial_positions()

	perform_dance(positions, instructions)
	fmt.Println(positions)

	test := []string {"a","b","c","d","e"}
	test_inst := []string {"s1","x3/4","pe/b"}
	perform_dance(test,test_inst)
	fmt.Println(test)
}
