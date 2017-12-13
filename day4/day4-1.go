package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"sort"
)

func check_data(f string) int {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	num_valid := 0
	for s.Scan() {
		line := s.Text()
		vals := strings.Fields(line)
		m := make(map[string]int)
		for _, v := range vals {
			// sort all the strings first, then we can just do the same uniqueness check

			// break it up into a slice first
			s_sl := strings.Split(v,"")
			// sort that slice!!
			sort.Strings(s_sl)
			// join it back up
			srt_s := strings.Join(s_sl,"")

			m[srt_s] = 0
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
