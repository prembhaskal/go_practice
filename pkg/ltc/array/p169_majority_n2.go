package array

func majorityElement(nums []int) int {
	return mooresVotingAlgo(nums)
}

func majorityElement1(nums []int) int {
	n := len(nums)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > n/2 {
			return v
		}
	}

	panic("should not happen")
}

func mooresVotingAlgo(nums []int) int {
	candidate := -1
	count := 0
	for _, v := range nums {
		if count == 0 { // discard prev array. if count reaches zero.
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
