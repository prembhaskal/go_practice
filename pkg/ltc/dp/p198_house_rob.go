package dp

func rob(nums []int) int {
	return robdpnoarray(nums)
}

func robdpnoarray(nums []int) int {
	n := len(nums)
	doubprev := nums[0]
	if n == 1 {
		return doubprev
	}
	prev := max(nums[0], nums[1])
	if n == 2 {
		return prev
	}

	// curr = max (nums[i] + doubprev, prev)
	var curr int
	for i := 2; i < n; i++ {
		curr = max(nums[i]+doubprev, prev)
		doubprev = prev
		prev = curr
	}

	return curr
}

func robdp(nums []int) int {
	// DP[i] => max money after visting 'i'th house.
	// DP[i] = Max( nums[i] + DP[i-2]  -> rob current house = current money + money obtained till 2 house back  )
	//              DP[i-1] -> don't rob current house, money obtained till previous house

	// base case
	// DP[0] = nums[0] -> only house rob it
	// DP[1] = max (nums[0], nums[1]) -> rob the one with max money.
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	if n == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	if n == 2 {
		return dp[1]
	}

	for i := 2; i < n; i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
