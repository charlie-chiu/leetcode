package pascals_triangle

func generate(numRows int) [][]int {
	var result = make([][]int, numRows+1)

	//row start from 1
	for row := 1; row <= numRows; row++ {
		for i := 0; i < row; i++ {
			// number on edge
			if i == 0 || i == row-1 {
				result[row] = append(result[row], 1)
			} else {
				sumAbove := result[row-1][i] + result[row-1][i-1]
				result[row] = append(result[row], sumAbove)
			}
		}
	}

	return result[1:]
}
