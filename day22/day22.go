package main

import (
	"os"
	"bufio"
	"fmt"
)

type Dir int
const (
	UP	= iota
	DOWN
	LEFT
	RIGHT
)

var infections = 0

func read_grid(f string) ([][]bool, int, int) {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	grid := make([][]bool, 0)
	for s.Scan() {
		line := s.Text()
		dline := make([]bool,len(line))
		for i, v := range(line) {
			if v == '#' {
				dline[i] = true
			} else {
				dline[i] = false
			}
		}
		grid = append(grid,dline)
	}

	gwidth := len(grid[0])
	gheight := len(grid)

	// build our full grid
	fullgrid := make([][]bool, 0)
	for i := 0; i < 100000; i++ {
		g := make([]bool,100000)
		fullgrid = append(fullgrid,g)
	}

	// add the config grid in the middle
	mx, my := 50000,50000
	sx, sy := (mx-(gheight/2)), (my-(gwidth/2))

	wx, wy := sx, sy
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fullgrid[wy][wx] = grid[i][j]
			wx++
		}
		wy++
		wx = sx
	}

	return fullgrid, mx, my

}

func turn_left(facing_dir Dir) Dir {
	switch(facing_dir) {
	case UP:
		return LEFT
	case DOWN:
		return RIGHT
	case LEFT:
		return DOWN
	case RIGHT:
		return UP
	}
	return UP
}

func turn_right(facing_dir Dir) Dir {
	switch(facing_dir) {
	case UP:
		return RIGHT
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case RIGHT:
		return DOWN
	}
	return UP
}

func iterate(grid [][]bool, facing_dir Dir, x int, y int) (Dir, int, int, bool) {


	var new_dir Dir
	new_infected := false
	if grid[y][x] == true {
		// infected	- turn right and make clean
		grid[y][x] = false
		new_dir = turn_right(facing_dir)
	} else {
		// not infected - infect and turn left
		grid[y][x] = true
		new_infected = true
		new_dir = turn_left(facing_dir)
	}

	switch(new_dir) {
	case UP:
		return new_dir,x,y-1,new_infected
	case DOWN:
		return new_dir,x,y+1,new_infected
	case LEFT:
		return new_dir,x-1,y,new_infected
	case RIGHT:
		return new_dir,x+1,y,new_infected
	}
	return new_dir,x,y,new_infected
}

func main() {

	grid, sx, sy := read_grid(os.Args[1])

	infections := 0
	var dir Dir = UP
	fmt.Println("Starting point value is: ",grid[sy][sx])
	fmt.Println("Left value is: ",grid[sy][sx-1])
	for i := 0; i < 10000; i++ {
		d, wx, wy, infect := iterate(grid,dir,sx,sy)
		dir = d
		sx = wx
		sy = wy
		if infect {
			infections++
		}
	}
	fmt.Println(infections)
}
