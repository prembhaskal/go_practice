package dp

import "fmt"

var (
	one  byte = '1'
	zero byte = '0'
)

// TODO - reduce space complexity.
func maximalSquare(matrix [][]byte) int {
	// return maximalSquare1(matrix)
	return maxSquareRec1(matrix)
}

var maxsqsize int

func maxSquareRec1(mat [][]byte) int {
	maxsqsize = 0
	m := len(mat)
	n := len(mat[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	maxSquareRec(mat, m-1, n-1, dp)
	return maxsqsize * maxsqsize
}

// returns the square size ending at vertex (r,c)
func maxSquareRec(mat [][]byte, r, c int, dp [][]int) int {
	// fmt.Printf("r: %d, c: %d\n", r, c)
	if r < 0 || c < 0 {
		return 0
	}
	if dp[r][c] != -1 {
		return dp[r][c]
	}

	// we traverse below squares too irrespective of whether current cell is '1' or '0'
	diag := maxSquareRec(mat, r-1, c-1, dp)
	vert := maxSquareRec(mat, r-1, c, dp)
	horz := maxSquareRec(mat, r, c-1, dp)
	var res int
	if mat[r][c] == one {
		res = 1 + min3(diag, vert, horz)
		if res > maxsqsize { // capture max size seen till now
			maxsqsize = res
		}
	}
	dp[r][c] = res
	// fmt.Printf("r:%d, c:%d, res: %d\n", r, c, dp[r][c])
	return dp[r][c]
}

func maximalSquare1(matrix [][]byte) int {
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
