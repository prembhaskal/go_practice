package binsearch

import (
	"fmt"
	"testing"
)

// 1(1) 2(1-2) 3(1-2-3) 4(1-2-3-4) 5(1-2-3-4-5)
func aTestBatchSizeSmallestDistancePair(t *testing.T) {
	nums := make([]int, 20)
	for k := 1; k <= 190; k++ {
		// k := 190
		batchnum := smallestDistancePair(nums, k)
		fmt.Printf("k: %d, batch: %d\n", k, batchnum)
		// fmt.Println()
	}
}

func TestSmallestDistancePair(t *testing.T) {
	// nums := []int{62, 100, 4}
	// k := 3
	nums := []int{9,10,7,10,6,1,5,4,9,8}
	k := 18

	dist := smallestDistancePair(nums, k)
	fmt.Printf("k: %d, kth dist: %d\n", k, dist)

	// 62, 100, 4
	// sorted 
	// 4, 62, 100
	// btch1: 58, 38  btch2: 96
}
