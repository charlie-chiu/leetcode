package valid_perfect_square

func isPerfectSquare(num int) bool {
	var root = num

	for root*root > num {
		root /= 2
	}

	for root*root < num {
		root += 1
	}

	if root*root == num {
		return true
	}
	return false
}
