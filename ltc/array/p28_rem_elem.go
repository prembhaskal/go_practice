package array

// https://leetcode.com/problems/remove-element/

// 3, 2, 2, 3

func removeElement(nums []int, val int) int {
	end := len(nums) - 1
	for i := 0; i <= end; {
		if nums[i] == val {
			nums[i] = nums[end]
			end--
		} else {
			i++
		}
	}

	return end + 1
}

func removeElement1(nums []int, val int) int {
	i := 0
	end := getLastNonMatchingIndex(nums, val, len(nums)-1)

	for ; i <= end; i++ {
		if nums[i] == val {
			// swap
			nums[i] = nums[end]
			nums[end] = val
		}

		end = getLastNonMatchingIndex(nums, val, end)
	}

	return end + 1
}

func getLastNonMatchingIndex(nums []int, val, currentPtr int) int {
	for ; currentPtr >= 0; currentPtr-- {
		if nums[currentPtr] != val {
			return currentPtr
		}
	}
	return -1
}
