package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func checksum_mod(i [][]int) int {
	sum := int(0)
	for _, b := range i {
		found_div := 0
		for i, c := range b {
			for j, d := range b {
				if i != j {
					if (c % d == 0) {
						found_div = c / d
					} else if (d % c == 0) {
						found_div = d / c
					}
				}
			}
		}
		sum += found_div
	}
	return sum
}

func readfile(f string) [][]int {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	var r [][]int = make([][]int,0)

	for s.Scan() {
		line := s.Text()
		vals := strings.Fields(line)
		var rr []int = make([]int,len(vals))
		for i, v := range vals {
			rr[i], _ = strconv.Atoi(string(v))
		}
		r = append(r,rr)
	}
	return r
}

func main() {
	datafile := os.Args[1]
	data := readfile(datafile)
	cs := checksum_mod(data)
	fmt.Println(cs)
}

