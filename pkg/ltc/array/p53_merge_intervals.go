package array

import "sort"

func merge(intervals [][]int) [][]int {
	// sort by start time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	overs := make([][]int, 0)
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		// if overlap
		if prev[1] >= curr[0] {
			newinter := []int{prev[0], max53(prev[1], curr[1])}
			prev = newinter
		} else {
			overs = append(overs, prev)
			prev = curr
		}
	}

	overs = append(overs, prev)

	return overs
}

func max53(a, b int) int {
	if a > b {
		return a
	}
	return b
}
