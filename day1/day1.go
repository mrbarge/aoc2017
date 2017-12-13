package main

// http://adventofcode.com/2017/day/1
// github.com/mrbarge/aoc2017

import "os"
import "strconv"
import "fmt"

func solve_captcha(s string) int {
	sl := len(s)
	sum := 0
	for i := 0; i < sl; i++ {
		d, _ := strconv.Atoi(string(s[i]))
		if s[i] == s[(i+1) % sl] {
			sum += d
		}
	}
	return sum
}

func main() {

	arg := os.Args[1]
	fmt.Println(solve_captcha(arg))
}