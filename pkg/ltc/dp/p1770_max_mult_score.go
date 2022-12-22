package dp

// TODO - add bottom up approach too, with optimal mem usage
func maximumScore(nums []int, multipliers []int) int {
	return maxScoreRecFunc(nums, multipliers)
}

var neginf = -3000000001

func maxScoreRecFunc(nums, mults []int) int {
	dp := make([][]int, len(mults))
	for i := 0; i < len(mults); i++ {
		dp[i] = make([]int, len(mults))
		for j := 0; j < len(mults); j++ {
			dp[i][j] = neginf
		}
	}
	return maxScoreRec(nums, mults, 0, 0, dp)
}

func maxScoreRec(nums, mults []int, idx, start int, dp [][]int) int {
	n := len(nums)
	m := len(mults)
	if start >= n || idx >= m {
		return 0
	}
	if dp[idx][start] != neginf {
		return dp[idx][start]
	}
	end := n - 1 - (idx - start)

	mult := mults[idx]
	// chose from start
	costStart := mult*nums[start] + maxScoreRec(nums, mults, idx+1, start+1, dp)

	// chose from end
	costEnd := mult*nums[end] + maxScoreRec(nums, mults, idx+1, start, dp)
	dp[idx][start] = max1770(costStart, costEnd)
	return dp[idx][start]
}

func max1770(a, b int) int {
	if a > b {
		return a
	}
	return b
}
