package main

import (
	"os"
	"bufio"
	"fmt"
	"errors"
	"regexp"
)

type Dir int
const (
	UP	= iota
	DOWN
	LEFT
	RIGHT
)

func read_map(f string) [][]rune {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	mapdata := make([][]rune, 0)

	i := 0
	for s.Scan() {

		line := s.Text()
		linemap := make([]rune,202)
		// initialise the line with spaces
		for i, _ := range(linemap) {
			linemap[i] = ' '
		}
		mapdata = append(mapdata,linemap)
		mapdata[i] = linemap

		for i, v := range(line) {
			linemap[i] = v
		}
		i++
	}

	return mapdata
}

func up(data [][]rune, x int, y int) (rune, error) {
	if (y == 0) {
		return ' ', errors.New("Cannot travel upwards")
	} else {
		return data[y-1][x], nil
	}
}

func down(data [][]rune, x int, y int) (rune, error) {
	if (y == len(data) - 1) {
		return ' ', errors.New("Cannot travel downwards")
	} else {
		return data[y+1][x], nil
	}
}

func left(data [][]rune, x int, y int) (rune, error) {
	if (x == 0) {
		return ' ', errors.New("Cannot travel left")
	} else {
		return data[y][x-1], nil
	}
}

func right(data [][]rune, x int, y int) (rune, error) {
	if (x == len(data[y]) - 1) {
		return ' ', errors.New("Cannot travel right")
	} else {
		return data[y][x+1], nil
	}
}

func is_end(data [][]rune, x int, y int, direction Dir) bool {
	r := data[y][x]

	var next_rune rune = ' '
	var next_err error = nil

	switch(direction) {
	case UP:
		if r == '+' {
			// we need to turn left or right
			l, lerr := left(data,x,y)
			r, rerr := right(data,x,y)
			if (lerr != nil && rerr != nil) {
				fmt.Println("Can't move any further ehadig down")
				return true
			} else if (l != ' ') {
				next_rune, next_err = left(data,x,y)
			} else if (r != ' ') {
				next_rune, next_err = right(data,x,y)
			}
		} else {
			// just keep on heading up if we can
			next_rune, next_err = up(data,x,y)
		}
	case DOWN:
		if r == '+' {
			// we need to turn left or right
			l, lerr := left(data,x,y)
			r, rerr := right(data,x,y)
			if (lerr != nil && rerr != nil) {
				fmt.Println("Can't move any further ehadig down")
				return true
			} else if (l != ' ') {
				next_rune, next_err = left(data,x,y)
			} else if (r != ' ') {
				next_rune, next_err = right(data,x,y)
			}
		} else {
			// just keep on heading up if we can
			next_rune, next_err = down(data,x,y)
		}
	case LEFT:
		if r == '+' {
			// we need to turn left or right
			u, uerr := up(data,x,y)
			d, derr := down(data,x,y)
			if (uerr != nil && derr != nil) {
				fmt.Println("Can't move any further ehadig left")
				return true
			} else if (u != ' ') {
				next_rune, next_err = up(data,x,y)
			} else if (d != ' ') {
				next_rune, next_err = down(data,x,y)
			}
		} else {
			// just keep on heading up if we can
			next_rune, next_err = left(data,x,y)
		}
	case RIGHT:
		if r == '+' {
			// we need to turn left or right
			u, uerr := up(data,x,y)
			d, derr := down(data,x,y)
			if (uerr != nil && derr != nil) {
				fmt.Println("Can't move any further ehadig left")
				return true
			} else if (u != ' ') {
				next_rune, next_err = up(data,x,y)
			} else if (d != ' ') {
				next_rune, next_err = down(data,x,y)
			}
		} else {
			// just keep on heading up if we can
			next_rune, next_err = right(data,x,y)
		}
	}

	if next_rune == ' ' || next_err != nil {
		return true
	} else {
		return false
	}
}

func travel(data [][]rune, x int, y int, direction Dir) (int, int, Dir) {
	r := data[y][x]
	rletter, _ := regexp.Compile(`[A-Z]`)
	lm := rletter.FindStringSubmatch(string(r))
	if lm != nil {
		fmt.Println("We passed a letter: ", string(r))
	}

	switch(direction) {
	case UP:
		if r == '+' {
			// we need to turn left or right
			l, lerr := left(data,x,y)
			r, rerr := right(data,x,y)
			if (lerr == nil && l != ' ') {
				return x-1,y,LEFT
			} else if (rerr == nil && r != ' ') {
				return x+1,y,RIGHT
			}
		} else {
			return x,y-1,UP
		}
	case DOWN:
		if r == '+' {
			// we need to turn left or right
			l, lerr := left(data,x,y)
			r, rerr := right(data,x,y)
			if (lerr == nil && l != ' ') {
				return x-1,y,LEFT
			} else if (rerr == nil && r != ' ') {
				return x+1,y,RIGHT
			}
		} else {
			return x,y+1,DOWN
		}
	case LEFT:
		if r == '+' {
			// we need to turn left or right
			u, uerr := up(data,x,y)
			d, derr := down(data,x,y)
			if (derr == nil && d != ' ') {
				return x,y+1,DOWN
			} else if (uerr == nil && u != ' ') {
				return x,y-1,UP
			}
		} else {
			return x-1,y,LEFT
		}
	case RIGHT:
		if r == '+' {
			// we need to turn left or right
			u, uerr := up(data,x,y)
			d, derr := down(data,x,y)
			if (derr == nil && d != ' ') {
				return x,y+1,DOWN
			} else if (uerr == nil && u != ' ') {
				return x,y-1,UP
			}
		} else {
			return x+1,y,RIGHT
		}
	}
	return 0,0,UP
}

func follow_map(data [][]rune) {

	x, y := 0, 0
	var direction Dir = DOWN

	// find starting position first
	for i, v := range(data[0]) {
		if v == '|' {
			x=i
			break
		}
	}

	steps := 0
	for done := false; !done; {

		if is_end(data,x,y,direction) {
			fmt.Println("we're done at: ", string(data[y][x]))
			done = true
		} else {
			x,y,direction = travel(data,x,y,direction)
			steps++
		}

	}
	fmt.Println(steps)
}

func main() {
	mapdata := read_map(os.Args[1])
	follow_map(mapdata)
}
