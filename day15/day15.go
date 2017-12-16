package main

import (
	"flag"
	"fmt"
	"testing"
)

/* Verify lower 16 bits match - straight stringcomp should work */

func lower_bits_match(a int64, b int64) bool {
	if (a & 0x0000FFFF) == (b & 0x0000FFFF) {
		return true
	} else {
		return false
	}
}

func gen_eval(gen_a_start int64, gen_b_start int64, gen_a_factor int64, gen_b_factor int64,
	iterations int64, divider int64) int64 {

	fmt.Printf("Generator A %d, Generator B %d\n", gen_a_start, gen_b_start)

	var i, genA, genB, matches int64 = 0, gen_a_start, gen_b_start, 0
	for i = 0; i < iterations; i++ {
		genA = (genA * gen_a_factor) % divider
		genB = (genB * gen_b_factor) % divider

		//fmt.Printf("A: %d, B: %d\n", genA, genB)
		if lower_bits_match(genA, genB) {
			matches++
		}
	}

	return matches

}

func TestLowerBitsMatch(t *testing.T) {
	if lower_bits_match(710, 711) {
		t.Fail()
	}
	if !lower_bits_match(245556042, 1431495499) {
		t.Fail()
	}
}

func main() {

	genAStartPtr := flag.Int64("genAStart", 722, "Generator A starting Value")
	genBStartPtr := flag.Int64("genBStart", 354, "Generator B starting Value")
	genAFactorPtr := flag.Int64("genAFactor", 16807, "Generator A factor")
	genBFactorPtr := flag.Int64("genBFactor", 48271, "Generator B factor")

	flag.Parse()
	m := gen_eval(*genAStartPtr, *genBStartPtr, *genAFactorPtr, *genBFactorPtr, 40000000, 2147483647)
	fmt.Println(m)
}

/*

Each of them to generate its next value, compares
the lowest 16 bits of both values, and keeps track of the number of times those parts
of the values match.

The generators both work on the same principle. To create its next value, a generator
will take the previous value it produced, multiply it by a factor (generator A uses
16807; generator B uses 48271), and then keep the remainder of dividing that resulting
product by 2147483647. That final remainder is the value it produces next.

To calculate each generator's first value, it instead uses a specific starting value as
its "previous value" (as listed in your puzzle input).

For example, suppose that for starting values, generator A uses 65, while generator B
uses 8921. Then, the first five pairs of generated values are:

--Gen. A--  --Gen. B--
   1092455   430625591
1181022009  1233683848
 245556042  1431495498
1744312007   137874439
1352636452   285222916
In binary, these pairs are (with generator A's value first in each pair):

00000000000100001010101101100111
00011001101010101101001100110111

01000110011001001111011100111001
01001001100010001000010110001000

00001110101000101110001101001010
01010101010100101110001101001010

01100111111110000001011011000111
00001000001101111100110000000111

01010000100111111001100000100100
00010001000000000010100000000100
Here, you can see that the lowest (here, rightmost) 16 bits of the third value match:
1110001101001010. Because of this one match, after processing these five pairs, the
judge would have added only 1 to its total.

To get a significant sample, the judge would like to consider 40 million pairs. (In
the example above, the judge would eventually find a total of 588 pairs that match
in their lowest 16 bits.)

After 40 million pairs, what is the judge's final count?

 */