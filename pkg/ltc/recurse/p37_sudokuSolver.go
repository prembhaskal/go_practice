package recurse

func solveSudoku(board [][]byte) {
	solveSudokuR(board)
}

// solve using back tracking
func solveSudokuR(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for n := 1; n < 10; n++ {
					if isValid(board, i, j, byte('0'+n)) {
						board[i][j] = byte('0' + n)
						// if this choice works, then we return
						// we don't need to check other cells in success case, as they would be already checked in below recursive call
						if solveSudokuR(board) {
							return true
						} else {
							// if this choice does not work, we reset and try next one.
							board[i][j] = '.'
						}
					}
				}
				// if none of the 9 choices work, we go back and try next choices
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, x, y int, ch byte) bool {
	for i := 0; i < 9; i++ {
		// check cols if other ch present
		if board[x][i] == ch {
			return false
		}

		// check rows
		if board[i][y] == ch {
			return false
		}
	}

	// find the block start row/col
	r := (x / 3) * 3
	c := (y / 3) * 3
	for i := r; i < r+3; i++ {
		for j := c; j < c+3; j++ {
			if board[i][j] == ch {
				return false
			}
		}
	}

	return true
}
