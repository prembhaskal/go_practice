package sort

import (
	"fmt"
	"testing"
)

func TestKSelect(t *testing.T) {
	// func findKthLargest(nums []int, k int) int {

	nums := []int{3, 4, 7, 1, 0, 9, 8}
	fmt.Printf("before test :%v\n", nums)
	k := 7
	numk := findKthLargest(nums, k)
	fmt.Printf("%d largest number is %d\n", k, numk)
}
