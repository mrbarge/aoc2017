package main

import "fmt"

type State int
const (
	A	= iota
	B
	C
	D
	E
	F
)

func iterate(s State, i int, tape []bool) (State, int) {
	new_idx := i
	new_state := s
	switch(s) {
	case A:
		if tape[i] {
			tape[i] = !tape[i]
			new_idx = i-1
			new_state = F
		} else {
			tape[i] = !tape[i]
			new_idx = i+1
			new_state = B
		}
	case B:
		if tape[i] {
			tape[i] = !tape[i]
			new_idx = i+1
			new_state = D
		} else {
			new_idx = i+1
			new_state = C
		}
	case C:
		if tape[i] {
			new_idx = i+1
			new_state = E
		} else {
			tape[i] = !tape[i]
			new_idx = i-1
			new_state = D
		}
	case D:
		if tape[i] {
			tape[i] = !tape[i]
			new_idx = i-1
			new_state = D
		} else {
			new_idx = i-1
			new_state = E
		}
	case E:
		if tape[i] {
			new_idx = i+1
			new_state = C
		} else {
			new_idx = i+1
			new_state = A
		}
	case F:
		if tape[i] {
			new_idx = i+1
			new_state = A
		} else {
			tape[i] = !tape[i]
			new_idx = i-1
			new_state = A
		}
	}

	return new_state, new_idx
}

func checksum(tape []bool) int {
	sum := 0
	for _, v := range(tape) {
		if v {
			sum++
		}
	}
	return sum
}

func main() {

	tape := make([]bool,500000)
	idx := 250000
	var s State = A
	for i := 0; i < 12794428; i++ {
		s, idx = iterate(s,idx,tape)
	}
	fmt.Println(checksum(tape))
}

