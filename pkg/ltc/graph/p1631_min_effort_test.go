package graph

import (
	"fmt"
	"testing"
)

func TestMinEffort(t *testing.T) {
	var heights [][]int

	heights = [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	minEffort := minimumEffortPath(heights)
	fmt.Printf("heights: %v, min effort: %d\n", heights, minEffort)

	heights = [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	minEffort = minimumEffortPath(heights)
	fmt.Printf("heights: %v, min effort: %d\n", heights, minEffort)

	heights = [][]int{{4, 3, 4, 10, 5, 5, 9, 2}, {10, 8, 2, 10, 9, 7, 5, 6}, {5, 8, 10, 10, 10, 7, 4, 2}, {5, 1, 3, 1, 1, 3, 1, 9}, {6, 4, 10, 6, 10, 9, 4, 6}}
	minEffort = minimumEffortPath(heights)
	fmt.Printf("heights: %v, min effort: %d\n", heights, minEffort)
}
