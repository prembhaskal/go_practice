package dp

// Kadane's algorithm.
func maxSubArray(nums []int) int {
	mxsum := -10001
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum > mxsum {
			mxsum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return mxsum
}
