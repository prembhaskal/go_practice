package array

// DP[i] = {max{DP[j]]} where j = i-1 to 0 and A[j] < A[i]}
// we use tail table to search in the smallest element larger than current element.

// tail table intuition
// tt[i] = x, it means 'x' is the smallest element using which we can have LIS of length 'i'
// whenever new number 'y' comes in, we check smallest number bigger than or equal to 'x', aka search for lower_bound
// tt is by definition sorted in increasing order (try on paper if you are not convinced)
// longest increasing subsequence
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tt := make([]int, 0)
	tt = append(tt, nums[0])
	ln := 1

	for i := 1; i < len(nums); i++ {
		last_elem := tt[ln-1]
		if nums[i] > last_elem {
			tt = append(tt, nums[i])
			ln++
		} else { // search element in tail table
			elem_idx := lower_bound(tt, nums[i])
			tt[elem_idx] = nums[i]
		}
	}

	return ln
}

// search for smallest idx, such that ar[idx] >= key
func lower_bound(ar []int, key int) int {
	low := 0
	high := len(ar) - 1
	ans := -1
	for low <= high {
		mid := (low + high) / 2
		if ar[mid] == key {
			return mid
		}
		if ar[mid] >= key {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
