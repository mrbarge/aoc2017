package main

import (
	"fmt"
	"os"
)

func parse_groups(data string) int {

	score := 0
	group_level := 0
	in_garbage := false
	in_cancel := false

	for i := 0; i < len(data); i++ {
		if (in_cancel) {
			in_cancel = false
			continue
		}

		if data[i] == '!' {
			in_cancel = true
			continue
		}

		if (!in_garbage) {
			if data[i] == '{' {
				group_level++
			} else if data[i] == '}' {
				score += group_level
				group_level--
			} else if data[i] == '<' {
				in_garbage = true
			}
		} else {
			if data[i] == '>' {
				in_garbage = false
			}
		}
	}
	return score

}

func main() {
	score := parse_groups(os.Args[1])
	fmt.Println(score)
}