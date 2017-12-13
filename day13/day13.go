package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"fmt"
)

func read_file(f string) [][]int{

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

	// build our data structure
	firewall := make([][]int,max_depth+1)
	for d, r := range(depths) {
		rl := make([]int,r)
		firewall[d] = rl
	}

	return firewall
}

func main() {

	firewall := read_file(os.Args[1])
	fmt.Println(firewall)
}
