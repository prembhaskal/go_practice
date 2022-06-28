package array

// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
// Return any array that satisfies this condition.
// TODO - write one variant with 1 pointer from back and 1 from front.
func sortArrayByParity(nums []int) []int {
	n := len(nums)

	od := 0

	// for ; od < n; od++ {
	// 	if nums[od]%2 == 1 {
	// 		break
	// 	}
	// }

	ev := 0
	// for ; ev < n; ev++ {
	// 	if nums[ev]%2 == 0 {
	// 		break
	// 	}
	// }

	// if od > ev {
	// 	return nums
	// }

	for ; ev < n && od < n; ev++ {
		if nums[ev]%2 == 1 {
			continue
		}
		if nums[od]%2 == 0 {
			od++
		}
		if od < n && nums[od]%2 == 1 && od < ev {
			tmp := nums[od]
			nums[od] = nums[ev]
			nums[ev] = tmp
			od++
		}
	}

	return nums
}
