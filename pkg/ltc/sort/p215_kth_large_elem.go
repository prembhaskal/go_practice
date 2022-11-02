package sort

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	return kselect(nums, 0, len(nums)-1, k)
}

func kselect(nums []int, start, end int, k int) int {
	fmt.Printf("kselect call, start: %d, end: %d\n", start, end)
	needIdx := len(nums) - k
	fmt.Printf("needIdx: %d\n", needIdx)
	idx := rand.Intn(end-start+1) + start
	fmt.Printf("idx is %d\n", idx)
	nums[idx], nums[end] = nums[end], nums[idx]

	paridx := partition215(nums, start, end)
	fmt.Printf("partition around: %d, paridx: %d, arr: %v\n", nums[paridx], paridx, nums)
	if paridx == needIdx {
		return nums[paridx]
	}
	if paridx > needIdx {
		return kselect(nums, start, paridx-1, k)
	}
	return kselect(nums, paridx, end, k)
}

func partition215(nums []int, start, end int) int {
	// fmt.Printf("start: %d, end: %d\n", start, end)
	key := nums[end] // last element is pivot
	fmt.Printf("partition key: %d\n", key)
	// fmt.Printf("partition pivot idx: %d, value: %d\n", end, key)

	i := start - 1
	j := start
	for ; j < end; j++ {
		// fmt.Printf("j:%d, nums[j]:%d\n", j, nums[j])
		if nums[j] <= key {
			i++
			// fmt.Printf("swap i: %d, j: %d\n", i, j)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = key, nums[i+1]
	return i + 1
}
