package heapr

import (
	"fmt"
	"testing"
)

func TestSkyline1(t *testing.T) {
	// [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	outline := getSkyline(buildings)
	fmt.Printf("outline is %v\n", outline)
}

func TestSkyline2(t *testing.T) {
	// [[0,2,3],[2,5,3]] input
	// [[0,3],[5,0]] output

	buildings := [][]int{
		{0, 2, 3},
		{2, 5, 3},
	}
	outline := getSkyline(buildings)
	fmt.Printf("outline is %v\n", outline)
}

func TestSkyline3(t *testing.T) {
	// [[0,3,3],[1,5,3],[2,4,3],[3,7,3]] input

	buildings := [][]int{
		{0, 3, 3},
		{1, 5, 3},
		{2, 4, 3},
		{3, 7, 3},
	}
	outline := getSkyline(buildings)
	fmt.Printf("outline is %v\n", outline)
}

func TestSkyline4(t *testing.T) {
	// [[1,5,3], [1,5,3], [1,5,3]] input

	buildings := [][]int{
		{1, 5, 3},
		{1, 5, 3},
		{1, 5, 3},
	}
	outline := getSkyline(buildings)
	fmt.Printf("outline is %v\n", outline)
}
