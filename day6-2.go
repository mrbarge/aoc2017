package main

import (
	"fmt"
	"errors"
	"strings"
)

func get_largest_block_idx(blocks []int) int {
	idx := 0
	largest := blocks[idx]
	for i, v := range(blocks) {
		if v > largest {
			idx = i
			largest = v
		}
	}
	return idx
}

func check_previous(patterns map[string]int, blocks string) (int, error) {
	if val, ok := patterns[blocks]; ok {
		return val, nil
	} else{
		return -1, errors.New("not found")
	}
}

func balance(blocks []int) int {
	done := false

	// do a map of block patterns to their starting cycle
	seen_patterns := make(map[string]int)

	num_blocks := len(blocks)
	cycles := 0
	cycle_diff := 0

	for !done {

		lb := get_largest_block_idx(blocks)
		nb := blocks[lb]
		// reset the largest
		blocks[lb] = 0
		// rebalance
		block_idx := (lb + 1) % num_blocks
		for i := 0; i < nb; i++ {
			blocks[block_idx]++
			block_idx = (block_idx + 1) % num_blocks
		}

		// lazy, copying https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string/37533144
		pattern := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(blocks)), ","), "[]")
		nc, err := check_previous(seen_patterns, pattern)
		if err == nil {
			cycle_diff = cycles - nc
			done = true
		} else {
			seen_patterns[pattern] = cycles
		}
		cycles++
	}
	return cycle_diff
}

func main() {
	data := []int{ 2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14 }
	//data := []int{0,2,7,0}
	cycles := balance(data)
	fmt.Println(cycles)
}
