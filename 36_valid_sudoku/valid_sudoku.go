package valid_sudoku

// 4ms, 72.92%
func isValidSudoku(board [][]byte) bool {
	return isColValid(board) && isRowValid(board) && isSubBoxValid(board)
}

func isRowValid(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		appeared := make(map[byte]bool)
		for col := 0; col < 9; col++ {

			elem := board[row][col]
			if _, ok := appeared[elem]; ok && elem != '.' {
				return false
			} else {
				appeared[elem] = true
			}
		}
	}

	return true
}

func isColValid(board [][]byte) bool {
	for col := 0; col < 9; col++ {
		appeared := make(map[byte]bool)
		for row := 0; row < 9; row++ {

			elem := board[row][col]
			if _, ok := appeared[elem]; ok && elem != '.' {
				return false
			} else {
				appeared[elem] = true
			}
		}
	}

	return true
}

func isSubBoxValid(board [][]byte) bool {
	//0,1,2 & 3,4,5 & 6,7,8
	//var row, col = 0, 0

	for col := 0; col < 9; col += 3 {
		for row := 0; row < 9; row += 3 {
			appeared := make(map[byte]bool)
			for subRow := row; subRow < row+3; subRow++ {
				for subCol := col; subCol < col+3; subCol++ {
					elem := board[subRow][subCol]
					if _, ok := appeared[elem]; ok && elem != '.' {
						return false
					} else {
						appeared[elem] = true
					}
				}
			}
		}
	}

	return true
}
