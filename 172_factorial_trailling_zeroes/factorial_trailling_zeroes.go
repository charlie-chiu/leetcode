package factorial_trailling_zeroes

import (
	"log"
	"math"
)

var solution = try2

// 4ms / 2.2MB
func try2(n int) int {
	var trailingZeroes int
	for n > 0 {
		/*
			625 / 5 = 125, we got 125 "5" in number 1 to 625
			x 5 x ... x 10 x ... x 15 x ... x 125 x ... x 625
			X 1 X ... X 2  x ... x 3  x ... x  5  x ... x 125
			repeat till no more 5
		*/
		n /= 5
		trailingZeroes += n
	}
	return trailingZeroes
}

// TLE : time limit exceeded
func try1(n int) int {
	var memory = make(map[int]int)

	var trailingZeroes int
	for i := 1; i <= n; i++ {
		curMultiplier := i

		zeroes, ok := memory[curMultiplier]
		if !ok {
			zeroes = 0
			for curMultiplier%5 == 0 {
				curMultiplier /= 5
				zeroes++
			}
			memory[curMultiplier] = zeroes
		}

		trailingZeroes += zeroes
	}

	return trailingZeroes
}

// overflow at 21!
func bruteForce(n int) int {
	//brute force

	// calculate factorial of n

	var factorial uint64 = 1
	for i := 1; i <= n; i++ {
		factorial *= uint64(i)

	}
	log.Printf("uint64 maximum:, %d", uint64(math.MaxUint64))
	log.Printf("%d! = %d", n, factorial)

	// count trailing zeroes
	var countOfZeroes int
	for factorial%10 == 0 {
		countOfZeroes++
		factorial /= 10
	}

	return countOfZeroes
}
