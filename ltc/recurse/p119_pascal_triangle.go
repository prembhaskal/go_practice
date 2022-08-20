package recurse

func getRow(rowIndex int) []int {
	// return getRowIter(rowIndex)
	return getRowRec(rowIndex, []int{1})
}

func getRowRec(rowIndex int, prev []int) []int {
	// fmt.Printf("rowIndex: %d, len(prev)=%d\n", rowIndex, len(prev))
	n := len(prev)
	if n == rowIndex+1 { // rowIndex is '0' based.
		return prev
	}
	curr := make([]int, n+1)
	for j := 0; j <= n; j++ {
		if j == 0 || j == n {
			curr[j] = 1
		} else {
			curr[j] = prev[j-1] + prev[j]
		}
	}

	return getRowRec(rowIndex, curr)
}

func getRowIter(rowIndex int) []int {
	prev := []int{1}
	for i := 1; i <= rowIndex; i++ {
		row := make([]int, rowIndex+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = prev[j-1] + prev[j]
			}

		}
		prev = row
	}

	return prev
}
