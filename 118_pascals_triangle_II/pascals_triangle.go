package pascals_triangle

func getRow(rowIndex int) []int {
	var result = []int{1}
	for row := 1; row <= rowIndex; row++ {
		for i := row - 1; i > 0; i-- {
			result[i] += result[i-1]
		}

		result = append(result, 1)
	}
	return result
}
