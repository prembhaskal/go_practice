package bit

func subsets(nums []int) [][]int {
	final := make([][]int, 0)
	return addsubsets(nums, make([]int, 0), 0, final)
}

func addsubsets(nums, curr []int, idx int, final [][]int) [][]int {
	if idx >= len(nums) {
		// fmt.Printf("final is %v\n", curr)
		final = append(final, curr)
		return final
	}

	// fmt.Printf("current is %v\n", curr)

	// fmt.Printf("choose current: %d\n", nums[idx])
	// choose current

	// create copy as append will usually modify input slice.
	dst := make([]int, len(curr))
	copy(dst, curr)
	ncurr := append(dst, nums[idx])
	final = addsubsets(nums, ncurr, idx+1, final)

	// fmt.Printf("don't choose current: %d\n", nums[idx])
	// don't choose current
	return addsubsets(nums, curr, idx+1, final)
}
