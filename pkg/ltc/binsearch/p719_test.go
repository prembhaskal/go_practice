package binsearch

import (
	"fmt"
	"testing"
)


func TestPairCount(t *testing.T) {
	nums := []int{1, 10, 10, 11, 15, 20}
	for i := 0; i <= 19; i++ {
		count := paircount(nums, i)
		fmt.Printf("distance pair count less than equal to %d is %d\n", i, count)
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
