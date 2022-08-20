package wc_278

import (
	"fmt"
	"testing"
)

func TestFindHighScore(t *testing.T) {
	runTest5981([]int{1, 1, 1})
	runTest5981([]int{1, 1})
	runTest5981([]int{0, 0, 0})
	runTest5981([]int{0, 0})
	runTest5981([]int{0, 0, 1, 0})
}

func runTest5981(arr []int) {
	fmt.Println()
	fmt.Printf("input: %v\n", arr)
	high := maxScoreIndices(arr)
	fmt.Printf("score: %v\n", high)
}
