package binsearch

import (
	"fmt"
	"testing"
)

func TestSplitPossible(t *testing.T) {
	nums := []int{7, 2, 5, 10, 8}
	subsum := 9
	m := 2
	isPos := checkMinSplitPossible(nums, subsum, m)
	fmt.Printf("for nums: %v, sum: %d, split: %d, min split: %d\n", nums, subsum, m, isPos)

	nums = []int{5,5,5,5}
	subsum = 4
	m = 4
	isPos = checkMinSplitPossible(nums, subsum, m)
	fmt.Printf("for nums: %v, sum: %d, split: %d, min split: %d\n", nums, subsum, m, isPos)

	nums = []int{1, 4, 4}
	subsum = 2
	m = 3
	isPos = checkMinSplitPossible(nums, subsum, m)
	fmt.Printf("for nums: %v, sum: %d, split: %d, min split: %d\n", nums, subsum, m, isPos)
}

func TestSplitArray(t *testing.T) {
	nums := []int{7, 2, 5, 10, 8}
	m := 2
	sum := splitArray(nums, m)
	fmt.Printf("split array nums: %v, split: %d, sum: %d\n", nums, m, sum)

	nums = []int{5,5,5,5}
	m = 4
	sum = splitArray(nums, m)
	fmt.Printf("split array nums: %v, split: %d, sum: %d\n", nums, m, sum)

	nums = []int{1, 4, 4}
	m = 3
	sum = splitArray(nums, m)
	fmt.Printf("split array nums: %v, split: %d, sum: %d\n", nums, m, sum)
}