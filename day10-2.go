package main

import (
	"fmt"
	"encoding/hex"
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

func main() {

	datastr := `157,222,1,2,177,254,0,228,159,140,249,187,255,51,76,30`
	lengths := []byte(datastr)
	lengths = append(lengths,17,31,73,47,23)
	d := process_data(lengths, 256, 64)
	hash := dense_hash(d,256)
	fmt.Println(hex.EncodeToString(hash))
}
