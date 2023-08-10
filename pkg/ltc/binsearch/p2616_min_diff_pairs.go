package binsearch

import "sort"

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	// need_pairs = p
	// return rec(nums, 0, 0, 0)

	low := 0
	high := nums[len(nums)-1] - nums[0]
	ans := 0
	for low <= high {
		mid := low + (high-low)/2
		pos := canformpairs(nums, mid, p)
		// fmt.Printf("mid: %d, pos: %d\n", mid, pos)
		if pos {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

func canformpairs(nums []int, max_diff int, pairs int) bool {
	total := 0
	for i := 0; i < len(nums)-1; {
		if (nums[i+1] - nums[i]) <= max_diff {
			i = i + 2
			total++
		} else {
			i = i + 1
		}
	}
	return total >= pairs
}

// DP times out, it runtime complexity is O(n * p) , even memory complexity is similar, but it can be optimized a bit.
// func dpbottomup(nums []int, p int) int {
//     // i i+1 i+2
//     dp := make([][]int, 0)

//     dp[i][p] = max(dp[i-2][p-1],  abs(nums[i]-nums[i-1]))
//     curr[p] = max(curr_2[p-1], abs(nums[i], nums[i-1]))

//     curr_2 = curr_1
//     curr_1 = curr

// }

// sorted
//  choose index i,i+1 as a pair,     f(i+2, max(curr_max, abs(a[i]-a[i+1])), total_pairs+1)
//  skip index   i, go for next one

// var need_pairs int

// const large_num int = 1000_000_001

// func rec(nums []int, idx int, curr_max int, pairs int) int {
// 	if idx >= len(nums) {
// 		if pairs == need_pairs {
// 			return curr_max
// 		}
// 		return large_num
// 	}

// 	choose := large_num
// 	if idx+1 < len(nums) {
// 		choose = rec(nums, idx+2, max(curr_max, abs(nums[idx]-nums[idx+1])), pairs+1)
// 	}

// 	not_choose := rec(nums, idx+1, curr_max, pairs)

// 	return min(choose, not_choose)
// }

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
