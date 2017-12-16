package main

import (
	"flag"
	"fmt"
)

func lower_bits_match(a int64, b int64) bool {
	if (a & 0x0000FFFF) == (b & 0x0000FFFF) {
		return true
	} else {
		return false
	}
}

func valid_gen_a_result(a int64) bool {
	return (a % 4 == 0)
}

func valid_gen_b_result(b int64) bool {
	return (b % 8 == 0)
}

func generator_a_next(a int64, factor int64, divider int64) int64 {

	for a = (a*factor) % divider; !valid_gen_a_result(a); a = (a*factor) % divider {
	}
	return a
}

func generator_b_next(b int64, factor int64, divider int64) int64 {
	for b = (b*factor) % divider; !valid_gen_b_result(b); b = (b*factor) % divider {
	}
	return b
}

func gen_eval(gen_a_start int64, gen_b_start int64, gen_a_factor int64, gen_b_factor int64,
	iterations int64, divider int64) int64 {

	var i, genA, genB, matches int64 = 0, gen_a_start, gen_b_start, 0
	for i = 0; i < iterations; i++ {
		genA = generator_a_next(genA, gen_a_factor, divider)
		genB = generator_b_next(genB, gen_b_factor, divider)

		if lower_bits_match(genA, genB) {
			matches++
		}
	}

	return matches

}

func main() {

	genAStartPtr := flag.Int64("genAStart", 722, "Generator A starting Value")
	genBStartPtr := flag.Int64("genBStart", 354, "Generator B starting Value")
	genAFactorPtr := flag.Int64("genAFactor", 16807, "Generator A factor")
	genBFactorPtr := flag.Int64("genBFactor", 48271, "Generator B factor")

	flag.Parse()
	m := gen_eval(*genAStartPtr, *genBStartPtr, *genAFactorPtr, *genBFactorPtr, 5000000, 2147483647)
	fmt.Println(m)
}
