package array

func setZeroes(matrix [][]int) {
	setZeroes2(matrix)
}

func setZeroes2(mat [][]int) {
	// use first row and col as map

	col0 := 1
	// (0,0) is only for row0,
	// use col0 for columnn.

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				if i == 0 && j == 0 {
					col0 = 0
					mat[0][0] = 0
				} else if j == 0 {
					col0 = 0
				} else {
					mat[i][0] = 0 // ith row
					mat[0][j] = 0 // jth column
				}
			}
		}
	}

	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[0]); j++ {
			if mat[0][j] == 0 || mat[i][0] == 0 {
				mat[i][j] = 0
			}
		}
	}

	if mat[0][0] == 0 {
		for j := 0; j < len(mat[0]); j++ {
			mat[0][j] = 0
		}
	}

	if col0 == 0 {
		for i := 0; i < len(mat); i++ {
			mat[i][0] = 0
		}
	}
}

func setZeroes1(matrix [][]int) {
	drows := make(map[int]bool)
	dcols := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				drows[i] = true
				dcols[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		if drows[i] {
			resetRow(matrix, i)
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		if dcols[j] {
			resetCol(matrix, j)
		}
	}
}

func resetCol(mat [][]int, col int) {
	for i := 0; i < len(mat); i++ {
		mat[i][col] = 0
	}
}

func resetRow(mat [][]int, row int) {
	for j := 0; j < len(mat[0]); j++ {
		mat[row][j] = 0
	}
}
