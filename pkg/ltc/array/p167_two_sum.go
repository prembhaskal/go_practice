package array

func twoSum(numbers []int, target int) []int {
	ar := numbers
	start := 0
	end := len(numbers) - 1

	for start < end {
		sum := ar[start] + ar[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			return []int{start + 1, end + 1}
		}
	}
	return nil
}
