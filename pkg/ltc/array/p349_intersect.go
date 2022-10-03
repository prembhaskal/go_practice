package array

func intersection(nums1 []int, nums2 []int) []int {
	return intersect_using_map(nums1, nums2)
}

func intersect_using_map(a, b []int) []int {
	allmap := make(map[int]bool)
	for _, num := range a {
		allmap[num] = true
	}
	all := make([]int, 0)
	for _, num := range b {
		if allmap[num] {
			all = append(all, num)
			allmap[num] = false
		}
	}

	return all
}
