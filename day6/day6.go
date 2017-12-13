package main

import (
	"fmt"
	"reflect"
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

func check_previous(patterns [][]int, blocks []int) bool {
	for _, v := range(patterns) {
		if reflect.DeepEqual(v,blocks) {
			return true
		}
	}
	return false
}

func balance(blocks []int) int {
	done := false

	var seen_patterns [][]int = make([][]int,0)
	num_blocks := len(blocks)
	cycles := 0

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

		if check_previous(seen_patterns, blocks) {
			done = true
		} else {
			pattern := append([]int(nil), blocks...)
			seen_patterns = append(seen_patterns, pattern)
		}
		cycles++
	}
	return cycles
}

func main() {
	data := []int{ 2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14 }
	//data := []int{0,2,7,0}
	cycles := balance(data)
	fmt.Println(cycles)
}
