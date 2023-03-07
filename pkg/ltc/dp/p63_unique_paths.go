package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	grid := make([][]int, 0)

	for i := 0; i < m; i++ {
		row := make([]int, n)
		grid = append(grid, row)
	}

	if obstacleGrid[0][0] == 0 {
		grid[0][0] = 1
	}
	// first row
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j-1] == 0 && obstacleGrid[0][j] == 0 {
			grid[0][j] = grid[0][j-1]
		}
	}

	// first col
	for i := 1; i < m; i++ {
		if obstacleGrid[i-1][0] == 0 && obstacleGrid[i][0] == 0 {
			grid[i][0] = grid[i-1][0]
		}
	}

	// remaining
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			top := grid[i-1][j]
			if obstacleGrid[i-1][j] == 1 {
				top = 0
			}
			left := grid[i][j-1]
			if obstacleGrid[i][j-1] == 1 {
				left = 0
			}

			if obstacleGrid[i][j] == 0 {
				grid[i][j] = top + left
			}
		}
	}

	return grid[m-1][n-1]
}
