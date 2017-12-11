package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func walk(x int, y int, steps int) int {
	if (x == 0 && y == 0) {
		return steps
	} else if (y > 0 && x > 0) {
		return walk(y-1,x-1,steps+1)
	} else if (y > 0 && x == 0) {
		return walk(y-1, x, steps+1)
	} else if (y < 0 && x == 0) {
		return walk(y+1, x, steps+1)
	} else if (y < 0 && x > 0) {
		return walk(y+1,x-1,steps+1)
	} else if (y > 0 && x < 0) {
		return walk(y-1,x+1,steps+1)
	} else if (y < 0 && x < 0) {
		return walk(y+1, x+1, steps+1)
	} else if (y == 0 && x > 0) {
		return walk(y+1, x-1, steps+1)
	} else if (y == 0 && x < 0) {
		return walk(y-1, x+1, steps+1)
	} else {
		fmt.Printf("help %d, %d\n",x,y)
		return 0
	}
}

func proc_file(f string) {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)
	s.Scan()
	data := strings.Split(s.Text(),",")

	x, y, xmax, xmin, ymax, ymin := 0, 0, 0 ,0, 0, 0
	for _, v := range(data) {
		switch(v) {
		case "n":
			y++
		case "ne":
			y++
			x++
		case "nw":
			y++
			x--
		case "s":
			y--
		case "se":
			y--
			x++
		case "sw":
			y--
			x--
		}

		if y > ymax {
			ymax = y
		}
		if x > xmax {
			xmax = x
		}
		if x < xmin {
			xmin = x
		}
		if x < ymin {
			ymin = y
		}
	}

	fmt.Printf("Our final coordinates are (%d,%d)\n",x,y)
	fmt.Printf("Our dimensions are X: %d, Y: %d\n", (xmax-xmin), (ymax-ymin))

	fmt.Printf("starting a walk\n")
	steps := walk(x,y,0)
	fmt.Printf("total steps: %d\n",steps)
}

func main() {
	datafile := os.Args[1]
	proc_file(datafile)
}
