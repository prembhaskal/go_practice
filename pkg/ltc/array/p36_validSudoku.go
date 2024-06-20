package array

// runtime: O(n^2) memory: O(n)
func isValidSudoku(board [][]byte) bool {
	freq := make([]int, 10)
	// check all rows
	for i := 0; i < 9; i++ {
		clearFreq(freq)
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch != '.' {
				idx := ch - '0'
				if freq[idx] > 0 {
					return false
				}
				freq[idx]++
			}
		}

	}

	// check all cols
	for i := 0; i < 9; i++ {
		clearFreq(freq)
		for j := 0; j < 9; j++ {
			ch := board[j][i]
			if ch != '.' {
				idx := ch - '0'
				if freq[idx] > 0 {
					return false
				}
				freq[idx]++
			}
		}
	}

	// check all blocks
	for bx := 0; bx < 3; bx++ {
		for by := 0; by < 3; by++ {
			clearFreq(freq)
			for i := bx * 3; i < bx*3+3; i++ {
				for j := by * 3; j < by*3+3; j++ {
					ch := board[i][j]
					if ch != '.' {
						idx := ch - '0'
						if freq[idx] > 0 {
							return false
						}
						freq[idx]++
					}
				}
			}
		}
	}

	return true
}

// not using built in clear, to avoid slowness due to type inference and stuff.
func clearFreq(freq []int) {
	for i := 0; i < len(freq); i++ {
		freq[i] = 0
	}
}

// complexity is O(n^4) but in practice runs quite fast, 
// because O(1) memory complexity no extra allocation
func isValidSudoku1(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !isValid(board, i, j) {
				return false
			}
		}
	}

	return true
}

func isValid(board [][]byte, x, y int) bool {
	ch := board[x][y]
	if ch == '.' {
		return true
	}
	board[x][y] = '.'                   // temporarily
	defer func() { board[x][y] = ch }() // set back before return

	for i := 0; i < 9; i++ {
		if board[x][i] == ch {
			return false
		}
		if board[i][y] == ch {
			return false
		}
	}

	// find the block
	bx := x / 3
	by := y / 3

	for i := bx * 3; i < bx*3+3; i++ {
		for j := by * 3; j < by*3+3; j++ {
			if board[i][j] == ch {
				return false
			}
		}
	}
	return true
}