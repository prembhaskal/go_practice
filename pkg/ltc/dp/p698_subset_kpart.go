package dp

func canPartitionKSubsets(nums []int, k int) bool {
	return canPrec(nums, k)
}

func canPrec(nums []int, k int) bool {
	totalsum := 0
	for _, num := range nums {
		totalsum += num
	}
	if totalsum%k != 0 {
		return false
	}
	target_sum := totalsum / k

	picked := make([]bool, len(nums))

	return canP(nums, picked, 0, k, 0, target_sum)
}

// accepted but runs too slow
// TODO - make it faster
func canP(nums []int, picked []bool, curr_idx, ss_count, curr_sum, target_sum int) bool {
	// if all subsets done, stop
	// subsets_count
	if ss_count == 0 {
		return true
	}

	// keep track of current subset sum
	if curr_sum > target_sum {
		return false
	}

	// if curr_subset is fine, then check for next subset
	// start from beginning for next index so curr_idx = 0
	if curr_sum == target_sum {
		return canP(nums, picked, 0, ss_count-1, 0, target_sum)
	}

	// if curr_sum < target_sum still going on, check with next numbers
	for i := curr_idx; i < len(nums); i++ {
		if !picked[i] {
			picked[i] = true // pick num at index = i
			possible := canP(nums, picked, i+1, ss_count, curr_sum+nums[i], target_sum)
			if possible {
				return true
			}
			picked[i] = false // unpick num at index = 1. to try with other combination.
		}
	}

	return false
}
