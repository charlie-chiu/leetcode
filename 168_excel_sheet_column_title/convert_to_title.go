package excel_sheet_column_title

import (
	"strings"
)

var solution = convertToTitle

func convertToTitle(n int) string {
	var result = []string{}

	for n > 0 {
		result = append(result, string(rune((n-1)%26+'A')))

		n = (n - 1) / 26
	}

	resultLength := len(result)
	for i := 0; i < resultLength/2; i++ {
		result[i], result[resultLength-1-i] = result[resultLength-1-i], result[i]
	}

	return strings.Join(result, "")
}
