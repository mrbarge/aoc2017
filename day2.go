package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func checksum(i [][]int) int {
	sum := int(0)
	for _, b := range i {
		largest := b[0]
		smallest := b[0]
		for _, c := range b {
			if c > largest {
				largest = c
			} else if c < smallest {
				smallest = c
			}
		}
		diff := largest - smallest
		sum += diff
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
	cs := checksum(data)
	fmt.Println(cs)
}
