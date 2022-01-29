package array

import (
	"fmt"
	"testing"
)

func TestDupZero(t *testing.T) {
	// arr := []int{1, 0 , 0 , 2, 3}
	// runTest(arr)

	runTest([]int{1})
	runTest([]int{1, 0})
	runTest([]int{0, 1})
	runTest([]int{1, 0, 0, 0})
	runTest([]int{1, 0, 0, 2, 3})
	runTest([]int{1, 0, 0, 0, 2, 3})
	runTest([]int{1, 0, 2, 3, 0, 4, 5, 0})
}

func runTest(arr []int) {
	fmt.Println()
	fmt.Printf("before: %v\n", arr)
	duplicateZeros(arr)
	fmt.Printf("after: %v\n", arr)
}
