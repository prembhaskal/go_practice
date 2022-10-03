package array

func intersect(nums1 []int, nums2 []int) []int {
	// return intersect_using_map2(nums1, nums2)
	return intersect_using_map3(nums1, nums2)
}

func intersect_using_map3(a, b []int) []int {
	m1 := make(map[int]int)
	for _, num := range a {
		m1[num]++
	}

	all := make([]int, 0)
	for _, num := range b {
		if m1[num] > 0 {
			all = append(all, num)
			m1[num]--
		}
	}

	return all
}

func intersect_using_map2(a, b []int) []int {
	m1 := make(map[int]int)
	for _, num := range a {
		m1[num]++
	}
	m2 := make(map[int]int)
	for _, num := range b {
		m2[num]++
	}

	all := make([]int, 0)
	for k, v := range m1 {
		for i := 0; i < min(v, m2[k]); i++ {
			all = append(all, k)
		}
	}

	return all
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
