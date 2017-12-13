package main

import (
	"os"
	"strconv"
	"fmt"
)

type Dir int
const (
	UP	Dir = 0
	DOWN Dir = 1
	LEFT Dir = 2
	RIGHT Dir = 3
)

func build_spiral(n int) int {
	xpos := 0
	ypos := 0
	xposMax := 0
	yposMax := 0
	xposMin := 0
	yposMin := 0

	d := Dir(RIGHT)
	for i := 1; i <= n; i++ {
		switch(d) {
		case UP:
			ypos++
			if ypos > yposMax {
				yposMax = ypos
				d = Dir(LEFT)
			}
		case DOWN:
			ypos--
			if ypos < yposMin {
				yposMin = ypos
				d = Dir(RIGHT)
			}
		case RIGHT:
			xpos++
			if xpos > xposMax {
				xposMax = xpos
				d = Dir(UP)
			}
		case LEFT:
			xpos--
			if xpos < xposMin {
				xposMin = xpos
				d = Dir(DOWN)
			}
		}
	}

	// xpos and ypos should now be our grid coord, turn them into absolute value in case neg
	if (xpos < 0) {
		xpos = -xpos
	}
	if (ypos < 0) {
		ypos = -ypos
	}

	// and that should be our manhattan distance..
	return (xpos + ypos) - 1
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	o := build_spiral(n)
	fmt.Println(o)
}