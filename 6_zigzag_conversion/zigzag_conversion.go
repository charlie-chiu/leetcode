package zigzag_conversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var currRow = 1
	var asc bool = true

	var result = make([][]rune, numRows+1)

	// rune
	for _, char := range s {
		result[currRow] = append(result[currRow], char)

		if asc {
			currRow++
		} else {
			currRow--
		}

		if currRow == 1 || currRow == numRows {
			asc = !asc
		}
	}

	// convert to string
	sb := strings.Builder{}
	for row := 1; row <= numRows; row++ {
		for _, char := range result[row] {
			sb.WriteRune(char)
		}
	}

	return sb.String()
}
