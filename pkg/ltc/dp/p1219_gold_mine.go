package dp

func getMaximumGold(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	val := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visit := make([][]bool, 0)
			for k := 0; k < m; k++ {
				visit = append(visit, make([]bool, n))
			}
			// visit[i][j] = true
			// fmt.Printf("start with i:%d, j:%d\n", i, j)
			val = max1219(val, getmax(visit, grid, i, j, m, n))
			// fmt.Printf("i: %d, j: %d, val is %d\n", i, j, val)
		}
	}

	return val
}

func getmax(visit [][]bool, grid [][]int, i, j, m, n int) int {
	// fmt.Printf("getmax i:%d, j:%d\n", i, j)

	if outofbound(i, j, m, n) {
		return 0
	}

	if visit[i][j] {
		return 0
	}

	if grid[i][j] == 0 {
		return 0 // stop case
	}

	visit[i][j] = true

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	val := 0
	for _, dir := range dirs {
		val = max1219(val, grid[i][j]+getmax(visit, grid, i+dir[0], j+dir[1], m, n))
		// fmt.Printf("i: %d, j: %d\n", ni, nj)
	}

	visit[i][j] = false
	return val
}

func outofbound(i, j, m, n int) bool {
	if i < 0 || i >= m {
		return true
	}
	if j < 0 || j >= n {
		return true
	}
	return false
}

func max1219(a, b int) int {
	if a > b {
		return a
	}
	return b
}
