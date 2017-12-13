package main

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

type Node struct {
	name string
	weight int
	children []*Node
}

func process_data(f string) {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	m := make(map[string]*Node,0)
	r_parent, _ := regexp.Compile(`(\w+)\s+\((\d+)\)\s+->\s+(.+)`)
	r_child, _ := regexp.Compile(`(\w+)\s+\((\d+)\)`)

	for s.Scan() {
		line := s.Text()

		res_parent := r_parent.FindStringSubmatch(line)
		res_child := r_child.FindStringSubmatch(line)
		if res_parent != nil {
			parent := res_parent[1]
			weight, _ := strconv.Atoi(res_parent[2])
			children := strings.Split(strings.Replace(res_parent[3]," ", "", -1), ",")

			if _, exists := m[parent]; !exists {
				n := &Node{
					name: parent,
					weight: weight,
					children: make([]*Node, 0),
				}
				m[parent] = n
			} else {
				m[parent].weight = weight
			}

			n := m[parent]
			for _, v := range(children) {
				if _, exists := m[v]; !exists {
					c := &Node{
						name: v,
						weight: 0,
						children: make([]*Node, 0),
					}
					m[v] = c
					n.children = append(n.children, c)
				} else {
					n.children = append(n.children, m[v])
				}
			}

		} else if res_child != nil {
			node := res_child[1]
			weight, _ := strconv.Atoi(res_child[2])

			if _, exists := m[node]; !exists {
				n := &Node{
					name: node,
					weight: weight,
					children: make([]*Node, 0),
				}
				m[node] = n
			} else {
				m[node].weight = weight
			}
		} else {
			fmt.Printf("check your data, chief: %s\n",line)
		}
	}

	// it's mapped, let's check that tree
	for _, v := range m {
		success := check_weight(v)
		if !success {
			//fmt.Printf("There's a problem with node %s\n", k)
		}
	}
}

func get_combined_weight(n *Node) int {
	if len(n.children) == 0 {
		return n.weight
	} else {
		sumweight := n.weight
		for _, c := range(n.children) {
			sumweight += get_combined_weight(c)
		}
		return sumweight
	}
}

func check_weight(n *Node) bool {
	if len(n.children) == 0 {
		return true
	}
	start_weight := get_combined_weight(n.children[0])
	for _, c := range(n.children) {

		wc := get_combined_weight(c)
		if start_weight != wc {
			fmt.Printf("Weight balance problem at root node %s (child weight %s (%d) mismatch other weight %s (%d))\n", n.name, n.children[0].name, start_weight, c.name, wc)
		}
	}
	return true
}

func main() {
	datafile := os.Args[1]
	process_data(datafile)
}
