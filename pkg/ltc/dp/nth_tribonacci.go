package dp

func tribonacci(n int) int {
	ni := make([]int, 3)
	ni[0] = 0
	ni[1] = 1
	ni[2] = 1
	if n < 3 {
		return ni[n]
	}
	for i := 3; i <= n; i++ {
		val := ni[0] + ni[1] + ni[2]
		ni[0] = ni[1]
		ni[1] = ni[2]
		ni[2] = val
	}
	return ni[2]
}
