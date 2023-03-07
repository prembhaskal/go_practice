package dp

func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	// rob first, then skip last OR skip first, then consider last
	return max213(rob1(nums[0:n-1]), rob1(nums[1:n]))
}

func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	mxrob := make([]int, n)
	mxskip := make([]int, n)

	mxrob[0] = nums[0]

	for i := 1; i < n; i++ {
		mxrob[i] = mxskip[i-1] + nums[i]
		mxskip[i] = max213(mxskip[i-1], mxrob[i-1])
	}

	return max213(mxrob[n-1], mxskip[n-1])
}

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}
