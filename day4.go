package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func check_data(f string) int {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	num_valid := 0
	for s.Scan() {
		line := s.Text()
		vals := strings.Fields(line)
		// chuck them all in a map, we can then compare map length to slice length
		m := make(map[string]int)
		for _, v := range vals {
			m[v] = 0
		}
		if len(m) == len(vals) {
			num_valid++
		}
	}
	return num_valid
}

func main() {
	datafile := os.Args[1]
	nv := check_data(datafile)
	fmt.Println(nv)
}
