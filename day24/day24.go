package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

type Port struct {
	side1 int
	side1Connected bool
	side2 int
	side2Collected bool
}

var max_weight = 0

func read_config(f string) []Port {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	ports := make([]Port,0)

	for s.Scan() {
		line := strings.Split(s.Text(),"/")
		s1, _ := strconv.Atoi(line[0])
		s2, _ := strconv.Atoi(line[1])

		p := Port{s1,false,s2,false }
		ports = append(ports,p)
	}

	return ports
}

func connect_next(ports []Port, w int, sum int) {
	found_next := false
	for i, v := range(ports) {
		if v.side1 == w || v.side2 == w {
			found_next = true
			next_ports := append(append([]Port{}, ports[:i]...), ports[i+1:]...)
			new_sum := sum + v.side1 + v.side2
			if v.side1 == w {
				connect_next(next_ports, v.side2, new_sum)
			} else {
				connect_next(next_ports, v.side1, new_sum)
			}
		} else {
			continue
		}
	}
	if !found_next && sum > max_weight {
		max_weight = sum
	}
}

func build_bridges(ports []Port) {
	starting_points := make([]int,0)
	for i := 0; i < len(ports); i++ {
		v := ports[i]
		if v.side1 == 0 || v.side2 == 0 {
			starting_points = append(starting_points, i)
		}
	}
	fmt.Println(starting_points)
	for _, v := range(starting_points) {
		remaining_ports := append(append([]Port{}, ports[:v]...),ports[v+1:]...)
		if ports[v].side1 == 0 {
			next_weight := ports[v].side2
			connect_next(remaining_ports,next_weight,next_weight)
		} else if ports[v].side2 == 0 {
			next_weight := ports[v].side1
			connect_next(remaining_ports,next_weight,next_weight)
		}
	}
}

func main() {
	ports := read_config(os.Args[1])
	build_bridges(ports)
	fmt.Println(max_weight)
}
