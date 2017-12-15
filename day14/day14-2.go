package main

import (
	"fmt"
	"strconv"
	"encoding/hex"
)

// from day 10
func reverse(data []int, starting_pos int, length int) {

	data_size := len(data)	 	// deduct one to represent final pos in array
	lpos := starting_pos

	var rev_data = make([]int,length)
	// build up a reversed version of the portion
	for i, j := 0, (length-1); i < length; i, j = i+1, j-1 {
		rev_data[j] = data[lpos]
		lpos = (lpos + 1) % data_size
	}

	// replace the original data with the reversed version
	lpos = starting_pos
	for i := 0; i < length; i++ {
		data[lpos] = rev_data[i]
		lpos = (lpos + 1) % data_size
	}
}

// from day 10
func process_data(lengths []byte, list_size int, rounds int) []int {

	// initialise data
	var data []int = make([]int,list_size)
	for i := 0; i < list_size; i++ {
		data[i] = i
	}

	skip_size := 0
	curr_pos := 0

	for x := 0; x < rounds; x++ {
		for _, l := range(lengths) {

			if int(l) > list_size {
				fmt.Printf("Error: length too long: %d\n",l)
				continue
			}

			// apply the reversal
			reverse(data, curr_pos, int(l))

			// move the pointer forward length + skipsize
			curr_pos = (curr_pos + int(l) + skip_size) % list_size

			// increment skip
			skip_size++
		}
	}

	return data
}

// from day 10
func dense_hash(data []int, list_size int) []byte {
	ret_data_size := list_size / 16
	var retdata []byte = make([]byte,ret_data_size)

	for i := 0; i<ret_data_size; i++ {
		sp := i*16
		x := byte(data[sp])
		for j := sp+1; j < (sp+16); j++ {
			x = x ^ byte(data[j])
		}
		retdata[i] = x
	}
	return retdata
}

// from day 10
func make_knot_hash(input string) string {
	lengths := []byte(input)
	lengths = append(lengths,17,31,73,47,23)
	d := process_data(lengths, 256, 64)
	hash := dense_hash(d,256)
	return hex.EncodeToString(hash)
}

// from: https://forum.golangbridge.org/t/hex-to-binary-function/4560/5
func hex_to_bin(in byte) string {
	var out []byte
	for i := 7; i >= 0; i-- {
		b := (in >> uint(i))
		out = append(out, (b%2)+48)
	}
	return string(out)
}

func check_region(used_grid [][]bool, region_grid [][]int, next_region int, x int, y int) bool {
	if used_grid[y][x] == true && region_grid[y][x] == 0 {
		region_grid[y][x] = next_region
		set_region(used_grid,region_grid,next_region,y,x)
		return true
	} else {
		return false
	}
}

// lazy lazy lazy
func get_region(region_grid [][]int, i int, j int) int {

	x, y := 0, 0

	// check itself
	if region_grid[i][j] > 0 {
		return region_grid[i][j]
	}

	// check north
	x, y = j, i-1
	if i == 0 {
		y = 0
	}
	if region_grid[y][x] > 0 {
		return region_grid[y][x]
	}

	// check west
	x, y = j-1, i
	if j == 0 {
		x = 0
	}
	if region_grid[y][x] > 0 {
		return region_grid[y][x]
	}

	// check south
	x, y = j, i+1
	if i+1 == len(region_grid) {
		y = i
	}
	if region_grid[y][x] > 0 {
		return region_grid[y][x]
	}

	// check east
	x, y = j+1, i
	if j+1 == len(region_grid[i]) {
		x = j
	}
	if region_grid[y][x] > 0 {
		return region_grid[y][x]
	}

	return 0
}

func set_region(used_grid [][]bool, region_grid [][]int, next_region int, i int, j int) bool {
	// check if a region exists anywhere adjacent, define & set one if not, return
	// indication if a new region is created
	x, y := 0, 0

	region_exists := get_region(region_grid,i,j)
	set_a_region := false
	if region_exists > 0 {
		region_grid[i][j] = region_exists
	} else {
		set_a_region = true
		region_grid[i][j] = next_region
	}

	// check north
	x, y = j, i-1
	if i == 0 {
		y = 0
	}
	check_region(used_grid, region_grid, region_grid[i][j], x, y)

	// check west
	x, y = j-1, i
	if j == 0 {
		x = 0
	}
	check_region(used_grid, region_grid, region_grid[i][j], x, y)

	// check south
	x, y = j, i+1
	if i+1 == len(region_grid) {
		y = i
	}
	check_region(used_grid, region_grid, region_grid[i][j], x, y)

	// check east
	x, y = j+1, i
	if j+1 == len(region_grid[i]) {
		x = j
	}
	check_region(used_grid, region_grid, region_grid[i][j], x, y)

	return set_a_region
}

func build_region_count(used_grid [][]bool, region_grid [][]int) int {
	next_region := 1
	for i, v := range(used_grid) {
		for j, _ := range(v) {
			if (used_grid[i][j] == true) {
				nr := set_region(used_grid, region_grid, next_region, i, j)
				if nr {
					next_region++
				}
			}
		}
		fmt.Println("")
	}
	return next_region-1
}

func count_regions(input string) int {
	used_grid := make([][]bool,128)
	region_grid := make([][]int,128)
	for i := 0; i < 128; i++ {
		lv := make([]bool,128)
		used_grid[i] = lv
		rv := make([]int, 128)
		region_grid[i] = rv
		input_str := input + "-" + strconv.Itoa(i)
		hash := make_knot_hash(input_str)
		hash_bytes, _ := hex.DecodeString(hash)

		for j, v := range(hash_bytes) {
			hash_byte_str := hex_to_bin(byte(v))
			for k, w := range(hash_byte_str) {
				if w == '1' {
					jpos := (len(hash_byte_str) * j) + k
					used_grid[i][jpos] = true
				}
			}
		}
	}

	return build_region_count(used_grid, region_grid)
}

func main() {
	input := "ffayrhll"
	regions := count_regions(input)
	fmt.Println(regions)

}
