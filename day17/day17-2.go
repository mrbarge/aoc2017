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

	r2 := r
	for i := 1; i <= 50000000; i++ {
		r2 = r2.Move(stepsize)
		rnew := ring.New(1)
		rnew.Value = i
		r2.Link(rnew)
		r2 = r2.Next()
	}

	r = r.Next()
	fmt.Println(r.Value)
}