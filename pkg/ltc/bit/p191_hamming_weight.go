package bit

func hammingWeight(num uint32) int {
	// return hamWeightAnd(num)
	return hamWeightDivide(num)
}

func hamWeightDivide(num uint32) int {
	cnt := 0
	for num > 0 {
		if num%2 == 1 {
			cnt++
		}
		num = num / 2
	}
	return cnt
}

func hamWeightAnd(num uint32) int {
	cnt := 0
	for num > 0 {
		cnt++
		num = num & (num - 1)
	}
	return cnt
}
