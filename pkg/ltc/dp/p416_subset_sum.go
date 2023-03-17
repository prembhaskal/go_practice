package dp

func canPartition(nums []int) bool {
	totalsum := 0
	for _, num := range nums {
		totalsum += num
	}
	if totalsum%2 != 0 {
		return false
	}
	need := totalsum / 2

	// amount
	dp := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		row := make([]int, need+1)
		dp = append(dp, row)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < need+1; j++ {
			dp[i][j] = -1
		}
	}

	return canp1(nums, 0, need, dp)
}

func canp1(nums []int, curr, amount int, dp [][]int) bool {
	if amount == 0 {
		return true
	}
	if amount < 0 {
		return false
	}

	if curr >= len(nums) {
		return false
	}

	if dp[curr][amount] != -1 {
		return dp[curr][amount] == 1 // add 1 if possible.
	}

	// choose curr
	choose := canp1(nums, curr+1, amount-nums[curr], dp)
	notchoose := canp1(nums, curr+1, amount, dp)
	if choose || notchoose {
		dp[curr][amount] = 1
	} else {
		dp[curr][amount] = 0
	}
	return dp[curr][amount] == 1
}
