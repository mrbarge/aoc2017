package main

import (
	"math"
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

func sum_neighbours(spiral [][]int, xpos int, ypos int) int {
	return spiral[xpos-1][ypos] + spiral [xpos-1][ypos-1] + spiral[xpos][ypos-1] +
		spiral[xpos+1][ypos-1] + spiral[xpos+1][ypos] + spiral[xpos+1][ypos+1] +
		spiral[xpos][ypos+1] + spiral[xpos-1][ypos+1]
}

func build_spiral(max_val int) int {

	// determine dimensions of the spiral
	spiralSize := int(math.Sqrt(float64(max_val)))

	spiral := make([][]int, spiralSize)
	for i := range spiral {
		spiral[i] = make([]int, spiralSize)
	}

	// start in the middle
	xpos := int(spiralSize/2)
	ypos := int(spiralSize/2)

	xposMax := xpos
	yposMax := ypos
	xposMin := xpos
	yposMin := ypos

	spiral[xpos][ypos] = 1

	d := Dir(RIGHT)
	for i := 1; i <= max_val; i++ {
		switch(d) {
		case UP:
			ypos++
			spiral[xpos][ypos] = sum_neighbours(spiral,xpos,ypos)
			if ypos > yposMax {
				yposMax = ypos
				d = Dir(LEFT)
			}
		case DOWN:
			ypos--
			spiral[xpos][ypos] = sum_neighbours(spiral,xpos,ypos)
			if ypos < yposMin {
				yposMin = ypos
				d = Dir(RIGHT)
			}
		case RIGHT:
			xpos++
			spiral[xpos][ypos] = sum_neighbours(spiral,xpos,ypos)
			if xpos > xposMax {
				xposMax = xpos
				d = Dir(UP)
			}
		case LEFT:
			xpos--
			spiral[xpos][ypos] = sum_neighbours(spiral,xpos,ypos)
			if xpos < xposMin {
				xposMin = xpos
				d = Dir(DOWN)
			}
		}
		if spiral[xpos][ypos] > max_val {
			return spiral[xpos][ypos]
		}
	}
	return -1
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	o := build_spiral(n)
	fmt.Println(o)
}
