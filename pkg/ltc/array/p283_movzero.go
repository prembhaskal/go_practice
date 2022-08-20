package array

import "fmt"

// https://leetcode.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	// nzptr := 0
	n := len(nums)

	// j := 0
	// for ; j < n; j++ {
	// 	if nums[j] == 0 {
	// 		break
	// 	}
	// }
	// zptr := j

	zptr := 0
	swp := 0
	for i := zptr + 1; i < n && zptr < n; i++ {
		if nums[zptr] != 0 {
			zptr++
			continue
		}
		if nums[i] != 0 {
			swp++
			nums[zptr] = nums[i]
			nums[i] = 0
			zptr++
		}
	}

	fmt.Printf("total swaps: %d\n", swp)
}
