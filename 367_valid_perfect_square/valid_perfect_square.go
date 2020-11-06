package valid_perfect_square

import "math"

func isPerfectSquare(num int) bool {
	root := math.Sqrt(float64(num))
	return math.Trunc(root) == root
}
