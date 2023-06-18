package array

func generate(numRows int) [][]int {
	pt := make([][]int, 0)
	pt = append(pt, []int{1})

	// 1
	// 1 1
	// 1 2 1
	// 1 3 3 1

	for i := 1; i < numRows; i++ {
		n := i + 1
		row := make([]int, n)
		row[0] = 1
		row[n-1] = 1

		for r := 1; r < n-1; r++ {
			row[r] = pt[i-1][r-1] + pt[i-1][r]
		}
		pt = append(pt, row)
	}
	return pt
}
