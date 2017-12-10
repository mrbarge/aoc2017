package main

import (
"fmt"
"strings"
"strconv"
)

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

func process_data(lengths []int, list_size int) []int {

	// initialise data
	var data []int = make([]int,list_size)
	for i := 0; i < list_size; i++ {
		data[i] = i
	}

	skip_size := 0
	curr_pos := 0

	for _, l := range(lengths) {

		if l > list_size {
			fmt.Printf("Error: length too long: %d\n",l)
			continue
		}

		// apply the reversal
		reverse(data, curr_pos, l)

		// move the pointer forward length + skipsize
		curr_pos = (curr_pos + l + skip_size) % list_size

		// increment skip
		skip_size++
	}

	return data
}

func main() {
	datastr := strings.Split("157,222,1,2,177,254,0,228,159,140,249,187,255,51,76,30",",")
	var lengths []int = make([]int,len(datastr))
	for i, v := range(datastr) {
		lengths[i], _ = strconv.Atoi(v)
	}
	const list_size int = 256
	d := process_data(lengths,list_size)
	fmt.Println(d)
}
