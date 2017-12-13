package main

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func process_data(f string) {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	m := make(map[string]string)
	w := make(map[string]int)

	r_parent, _ := regexp.Compile(`(\w+)\s+\((\d+)\)\s+->\s+(.+)`)
	r_child, _ := regexp.Compile(`(\w+)\s+\((\d+)\)`)

	for s.Scan() {
		line := s.Text()

		res_parent := r_parent.FindStringSubmatch(line)
		res_child := r_child.FindStringSubmatch(line)
		if res_parent != nil {
			parent := res_parent[1]
			weight := res_parent[2]
			children := strings.Split(strings.Replace(res_parent[3]," ", "", -1), ",")
			for _, v := range(children) {
				m[v] = parent
				w[parent], _ = strconv.Atoi(weight)
			}
		} else if res_child != nil {
			node := res_child[1]
			weight := res_child[2]
			w[node], _ = strconv.Atoi(weight)
		} else {
			fmt.Printf("check your data, chief: %s\n",line)
		}
	}

	// our base should be the one node with a weight but no parent
	for k, _ := range(w) {
		if _, exists := m[k]; !exists {
			fmt.Printf("Our base is %s\n",k)
		}
	}
}

func main() {
	datafile := os.Args[1]
	process_data(datafile)
}
