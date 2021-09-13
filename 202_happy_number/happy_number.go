package happy_number

import (
	"sort"
)

func isHappy(n int) bool {
	var isShowed = make(map[int]bool)
	var newN int

	for n != 1 {
		newN = 0

		if _, ok := isShowed[n]; ok {
			return false
		}

		for _, digit := range getDigits(n) {
			newN += digit * digit
		}
		isShowed[n] = true
		n = newN
	}

	return true
}

func getDigits(n int) (digits []int) {
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	sort.Ints(digits)
	return
}
