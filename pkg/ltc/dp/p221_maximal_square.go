package dp

import "fmt"

var (
	one  byte = '1'
	zero byte = '0'
)

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	msq := make([][]int, rows)
	horzones := make([][]int, rows)
	vertones := make([][]int, rows)
	for i := 0; i < rows; i++ {
		msq[i] = make([]int, cols)
		horzones[i] = make([]int, cols)
		vertones[i] = make([]int, cols)
	}
	maxsize := 0
	// initialize base cases
	if matrix[0][0] == one {
		msq[0][0] = 1
		maxsize = 1
		horzones[0][0] = 1
		vertones[0][0] = 1
	}

	// first row
	for i := 1; i < cols; i++ {
		if matrix[0][i] == one {
			maxsize = 1
			horzones[0][i] = horzones[0][i-1] + 1
			vertones[0][i] = 1
			msq[0][i] = 1
			// fmt.Printf("first row, col: %d, horz: %v\n", i, horzones)
		}
	}

	// first column
	for i := 1; i < rows; i++ {
		if matrix[i][0] == one {
			maxsize = 1
			vertones[i][0] = vertones[i-1][0] + 1
			horzones[i][0] = 1
			msq[i][0] = 1
		}
	}

	// fmt.Printf("msq: %v, horz:%d, vert:%v\n", msq, horzones, vertones)
	// fmt.Println("msq")
	// printmatrix(msq)

	// fmt.Println("horzones")
	// printmatrix(horzones)

	// fmt.Println("vertones")
	// printmatrix(vertones)

	// remaining

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == one {
				horzones[i][j] = horzones[i][j-1] + 1
				vertones[i][j] = vertones[i-1][j] + 1
				msq[i][j] = min3(msq[i-1][j-1]+1, horzones[i][j], vertones[i][j])
				//   fmt.Printf("i:%d, j:%d, msq: %v, hor:%d, vert:%v\n", i, j, msq, horzones, vertones)
				if msq[i][j] > maxsize {
					maxsize = msq[i][j]
				}
			}
		}
	}

	return maxsize * maxsize
}

func printmatrix(mtx [][]int) {
	for i := 0; i < len(mtx); i++ {
		fmt.Printf("%v\n", mtx[i])
	}
}

func min3(a, b, c int) int {
	return min221(min221(a, b), c)
}

func min221(a, b int) int {
	if a < b {
		return a
	}
	return b
}
