package sort

// https://leetcode.com/problems/merge-sorted-array/submissions/

// nums1 has m + n, nums2 has n
// merge in new array and update in original array.
func merge1(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)

	// 1, 3, 5, 6, 0 ,0 , 0
	// 2, 4, 5
	// 1, 2, 3, 4, 5, 5, 6

	var i, j, k int // j to track nums1, k to track nums2
	for ; i < m+n; i++ {
		if j >= m {
			if k < n {
				nums3[i] = nums2[k]
				k++
			} else {
				panic("should not happen")
			}
		} else if k >= n {
			if j < m {
				nums3[i] = nums1[j]
				j++
			} else {
				panic("should not happend")
			}
		} else {
			if nums1[j] < nums2[k] {
				nums3[i] = nums1[j]
				j++
			} else if k < n {
				nums3[i] = nums2[k]
				k++
			}
		}
	}

	copy(nums1, nums3)
}

// nums1 has m + n, nums2 has n
// merge in new array and update in original array.
func merge2(nums1 []int, m int, nums2 []int, n int) {
	// 1, 3, 5, 6, 0 ,0 , 0
	// 2, 4, 5
	// 1, 2, 3, 4, 5, 5, 6

	i := m + n - 1
	j := m - 1
	k := n - 1
	for ; i >= 0; i-- {
		if j < 0 {
			if k >= 0 {
				nums1[i] = nums2[k]
				k--
			} else {
				panic("should not happen")
			}
		} else if k < 0 {
			if j >= 0 {
				nums1[i] = nums1[j]
				j--
			} else {
				panic("should not happen")
			}
		} else {
			if nums1[j] >= nums2[k] {
				nums1[i] = nums1[j]
				j--
			} else {
				nums1[i] = nums2[k]
				k--
			}
		}
	}
}

// nums1 has m + n, nums2 has n
// merge in new array and update in original array.
// merged boundary conditions into 1 if else block.
func merge3(nums1 []int, m int, nums2 []int, n int) {
	// 1, 3, 5, 6, 0 ,0 , 0
	// 2, 4, 5
	// 1, 2, 3, 4, 5, 5, 6
	i := m + n - 1
	j := m - 1
	k := n - 1
	for ; i >= 0; i-- {
		if k < 0 || j >= 0 && nums1[j] >= nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
	}
}
