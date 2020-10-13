package ugly_numbers

//from https://www.geeksforgeeks.org/ugly-numbers/

func isUgly(n int) bool {
	if n == 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}

func getNthUglyNumber(n int) int {
	var count int
	for i := 1; ; i++ {
		if isUgly(i) {
			count++
			if count == n {
				return i
			}
		}
	}
}
