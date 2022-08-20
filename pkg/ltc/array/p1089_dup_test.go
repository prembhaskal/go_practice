package array

import (
	"fmt"
	"testing"
)

func TestDupZero(t *testing.T) {
	// arr := []int{1, 0 , 0 , 2, 3}
	// runTestdupzero(arr)

	runTestdupzero([]int{1})
	runTestdupzero([]int{1, 0})
	runTestdupzero([]int{0, 1})
	runTestdupzero([]int{1, 0, 0, 0})
	runTestdupzero([]int{1, 0, 0, 2, 3})
	runTestdupzero([]int{1, 0, 0, 0, 2, 3})
	runTestdupzero([]int{1, 0, 2, 3, 0, 4, 5, 0})
}

func runTestdupzero(arr []int) {
	fmt.Println()
	fmt.Printf("before: %v\n", arr)
	duplicateZeros(arr)
	fmt.Printf("after: %v\n", arr)
}
