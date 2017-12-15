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

func count_used(input string) int {
	used := 0
	for i := 0; i < 128; i++ {
		input_str := input + "-" + strconv.Itoa(i)
		hash := make_knot_hash(input_str)
		hash_bytes, _ := hex.DecodeString(hash)
		for _, v := range(hash_bytes) {
			hash_byte_str := hex_to_bin(byte(v))
			for _, w := range(hash_byte_str) {
				if w == '1' {
					used++
				}
			}
		}
	}
	return used
}

func main() {
	input := "ffayrhll"
	used := count_used(input)
	fmt.Println(used)
}
