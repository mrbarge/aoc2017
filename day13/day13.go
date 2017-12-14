package main

import (
	"bufio"
	"regexp"
	"strconv"
	"fmt"
	"strings"
	"os"
)

func start_scanner(firewall [][]bool) {
	for i, v := range(firewall) {
		// check if there is a layer at that depth
		if len(v) > 0 {
			firewall[i][0] = true
		}
	}
}

func iterate_scanner(firewall [][]bool, direction []bool) {
	for i, v := range(firewall) {
		// ignore empty-range depths
		if len(v) == 0 {
			continue
		}

		scanner_pos := -1
		for j, w := range(v) {
			if w == true {
				scanner_pos = j
				break
			}
		}

		if scanner_pos < 0 {
			fmt.Printf("Scanner was at an invalid position, depth %d\n", i)
			continue
		}

		// move scanner on
		v[scanner_pos] = false
		// check direction, false=down, true=up
		if direction[i] {
			if scanner_pos == 0 {
				direction[i] = false
				scanner_pos++
			} else {
				scanner_pos--
			}
		} else {
			if scanner_pos == (len(v)-1) {
				direction[i] = true
				scanner_pos--
			} else {
				scanner_pos++
			}
		}
		v[scanner_pos] = true
	}
}

func read_file(f string) [][]bool{

	//file, _ := os.Open(f)
	s := bufio.NewScanner(strings.NewReader(f))

	r, _ := regexp.Compile(`(\d+): (\d+)`)

	max_depth := 0
	depths := make(map[int]int,0)

	for s.Scan() {
		line := s.Text()
		lm := r.FindStringSubmatch(line)
		if lm != nil {
			d, _ := strconv.Atoi(lm[1])
			r, _ := strconv.Atoi(lm[2])
			if d > max_depth {
				max_depth = d
			}
			depths[d] = r
		}
	}

	// build our data structure. true in nested dict indicates the scanner is scanning that pos
	firewall := make([][]bool,max_depth+1)
	for d, r := range(depths) {
		rl := make([]bool,r)
		firewall[d] = rl
	}

	return firewall
}

func main_loop(firewall [][]bool, delay int) bool {
	start_scanner(firewall)

	fw_direction := make([]bool,len(firewall))

	for i := 0; i < delay; i++ {
		iterate_scanner(firewall, fw_direction)
	}

	for i := 0; i < len(firewall); i++ {
		layer := firewall[i]
		if len(layer) > 0 && layer[0] == true {
			return false
		}
		iterate_scanner(firewall, fw_direction)
	}

	return true
}

func main() {

	input := ""

	delay := 0
	for done := false; !done {
		firewall := read_file(input)
		if main_loop(firewall, delay) {
			fmt.Printf("Success at picosecond %d\n", i)
			done = true
		}
		delay++
	}
}


package main

import (
"os"
"bufio"
"regexp"
"strconv"
"fmt"
)


func start_scanner(firewall [][]bool) {
	for i, v := range(firewall) {
		// check if there is a layer at that depth
		if len(v) > 0 {
			firewall[i][0] = true
		}
	}
}

func iterate_scanner(firewall [][]bool, direction []bool) {
	for i, v := range(firewall) {
		// ignore empty-range depths
		if len(v) == 0 {
			continue
		}

		scanner_pos := -1
		for j, w := range(v) {
			if w == true {
				scanner_pos = j
				break
			}
		}

		if scanner_pos < 0 {
			fmt.Printf("Scanner was at an invalid position, depth %d\n", i)
			continue
		}

		// move scanner on
		v[scanner_pos] = false
		// check direction, false=down, true=up
		if direction[i] {
			if scanner_pos == 0 {
				direction[i] = false
				scanner_pos++
			} else {
				scanner_pos--
			}
		} else {
			if scanner_pos == (len(v)-1) {
				direction[i] = true
				scanner_pos--
			} else {
				scanner_pos++
			}
		}
		v[scanner_pos] = true
	}
}

func read_file(f string) [][]bool{

	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	r, _ := regexp.Compile(`(\d+): (\d+)`)

	max_depth := 0
	depths := make(map[int]int,0)

	for s.Scan() {
		line := s.Text()
		lm := r.FindStringSubmatch(line)
		if lm != nil {
			d, _ := strconv.Atoi(lm[1])
			r, _ := strconv.Atoi(lm[2])
			if d > max_depth {
				max_depth = d
			}
			depths[d] = r
		}
	}

	// build our data structure. true in nested dict indicates the scanner is scanning that pos
	firewall := make([][]bool,max_depth+1)
	for d, r := range(depths) {
		rl := make([]bool,r)
		firewall[d] = rl
	}

	return firewall
}

func main_loop(firewall [][]bool) int {
	start_scanner(firewall)

	severity := 0
	catches := make([]int,0)
	fw_direction := make([]bool,len(firewall))

	for starting_pos := 0; starting_pos < len(firewall); starting_pos++ {
		layer := firewall[starting_pos]
		if len(layer) > 0 && layer[0] == true {
			fmt.Printf("caught at depth %d\n",starting_pos)
			catches = append(catches,starting_pos)
		}
		iterate_scanner(firewall, fw_direction)
	}

	for _, v := range(catches) {
		severity += (v * len(firewall[v]))
	}

	return severity
}

func main() {

	firewall := read_file(os.Args[1])

	severity := main_loop(firewall)
	fmt.Println(severity)
}
