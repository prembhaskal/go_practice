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
	prev := max1218(nums[0], nums[1])
	if n == 2 {
		return prev
	}

	// curr = max (nums[i] + doubprev, prev)
	var curr int
	for i := 2; i < n; i++ {
		curr = max1218(nums[i]+doubprev, prev)
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
	dp[1] = max1218(nums[0], nums[1])
	if n == 2 {
		return dp[1]
	}

	for i := 2; i < n; i++ {
		dp[i] = max1218(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



func recmemc(nums []int) int {
    mem := make([]int, len(nums))
    for i := 0; i < len(mem); i++ {
        mem[i] = -1
    }
    return recmem(nums, 0, mem)
}

// Approach - at each idx, check max answer obtained by either chosing or not chosing
// if we have already processed index at 'next' , return same answer as nothing else has changed.
func recmem(nums []int, next int, mem []int) int {
    if next >= len(nums) {
        return 0
    }
    if mem[next] != -1 {
        return mem[next]
    }

    // chose next, means next+1 cannot be chosen
    chose := recmem(nums, next + 2, mem) + nums[next]

    // not chose next, means next+1 can be chosen
    notchose := recmem(nums, next + 1, mem)

    mem[next] = max(chose, notchose)
    return mem[next]
}

// Approach - at each idx, check max answer obtained by either chosing or not chosing
func rec(nums []int, next int) int {
    if next >= len(nums) {
        return 0
    }

    // chose next, means next+1 cannot be chosen
    chose := rec(nums, next + 2) + nums[next]

    // not chose next, means next+1 can be chosen
    notchose := rec(nums, next + 1)

    return max(chose, notchose)
}