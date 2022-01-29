package array

import (
	"fmt"
	"testing"
)

func TestRemoveDup(t *testing.T) {
	runTest([]int{1, 1, 2})
	runTest([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}

func runTest(arr []int) {
	fmt.Println()
	fmt.Printf("before: %v\n", arr)
	cnt := removeDuplicates(arr)
	fmt.Printf("after: %v\n", arr)
	fmt.Printf("count: %d\n", cnt)
}
