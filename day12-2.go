package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"fmt"
)

var progs = make(map[int]map[int]bool)

func add_relationship(prog1 int, prog2 int) {
	if _, exists := progs[prog1]; !exists {
		var l = make(map[int]bool, 0)
		progs[prog1] = l
	}
	if _, exists := progs[prog2]; !exists {
		var l = make(map[int]bool, 0)
		progs[prog2] = l
	}
	progs[prog1][prog2] = true
	progs[prog2][prog1] = true
}

func read_file(f string) {

	file, _ := os.Open(f)
	s := bufio.NewScanner(file)


	r, _ := regexp.Compile(`(\d+) <-> (.+)`)

	for s.Scan() {
		line := s.Text()

		lm := r.FindStringSubmatch(line)
		if lm != nil {
			program, _ := strconv.Atoi(lm[1])
			prog_group := strings.Split(strings.Replace(lm[2]," ", "", -1), ",")
			for _, v := range(prog_group) {
				prog_link, _ := strconv.Atoi(v)
				add_relationship(program, prog_link)
			}
		}
	}
}

func count_links(prog int, seen map[int]bool) {
	if _, exists := seen[prog]; exists {
		return
	} else {
		seen[prog] = true
	}

	for k, _ := range(progs[prog]) {
		count_links(k, seen)
	}

	return
}

func main() {
	read_file(os.Args[1])
	seen := make(map[int]bool,0)

	groups := 0
	for k, _ := range(progs) {
		if _, exists := seen[k]; !exists {
			count_links(k, seen)
			groups++
		}
	}
	fmt.Println(groups)
}
