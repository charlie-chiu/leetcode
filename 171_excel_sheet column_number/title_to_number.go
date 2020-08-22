package excel_sheet_column_number

var solution = titleToNumber

func titleToNumber(s string) int {
	// get number from ascii

	const numOfAlphabets = 26
	var answer int32

	for i, char := range s {
		exponentiation := len(s) - 1 - i
		charNumber := char - 64
		for j := 0; j < exponentiation; j++ {
			charNumber *= numOfAlphabets
		}
		answer += charNumber
	}

	return int(answer)
}
