package array

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	runTestMZ([]int{1, 0, 5, 4, 0, 1})
	runTestMZ([]int{0, 0, 0, 0, 1, 2, 3})
	runTestMZ([]int{1, 2, 3, 5, 6, 7})
}

func runTestMZ(nums []int) {
	fmt.Println()
	fmt.Printf("before: %v\n", nums)
	moveZeroes(nums)
	fmt.Printf("after: %v\n", nums)
	fmt.Println()
}
