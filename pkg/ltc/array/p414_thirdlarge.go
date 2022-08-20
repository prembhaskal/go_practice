package array

func thirdMax(nums []int) int {
	infin := 1 << 32
	infin = -infin

	one, two, three := infin, infin, infin

	for i := 0; i < len(nums); i++ {
		if nums[i] >= one {
			if nums[i] != one {
				three = two
				two = one
				one = nums[i]
			}
		} else if nums[i] >= two {
			if nums[i] != two {
				three = two
				two = nums[i]
			}
		} else if nums[i] > three {
			three = nums[i]
		}
	}

	if three != infin {
		return three
	}
	return one
}
