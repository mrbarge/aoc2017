package main

import (
	"strconv"
	"os"
	"container/ring"
	"fmt"
)

func main() {

	stepsize, _ := strconv.Atoi(os.Args[1])

	// Initialise the ring
	r := ring.New(1)
	r.Value = 0

	for i := 1; i <= 2017; i++ {
		r = r.Move(stepsize)
		rnew := ring.New(1)
		rnew.Value = i
		r.Link(rnew)
		r = r.Next()
	}

	// Now shift one more to get our element after 2017
	r = r.Next()
	fmt.Println(r.Value)
}