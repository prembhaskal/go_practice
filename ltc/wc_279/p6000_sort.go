package wc_279

import (
	"sort"
)

// https://leetcode.com/contest/weekly-contest-279/problems/sort-even-and-odd-indices-independently/
func sortEvenOdd(nums []int) []int {
	n := len(nums)
	even := make([]int, 0, n/2+1)
	odd := make([]int, 0, n/2+1)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}

	sort.Sort(sort.IntSlice(even))
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			ans[i] = even[i/2]
		} else {
			ans[i] = odd[(i-1)/2]
		}
	}
	return ans
}
