package array

// https://leetcode.com/problems/valid-mountain-array/

// TODO - optimize further
func validMountainArray1(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	dec := false
	i := 1
	for i < n {
		if arr[i] > arr[i-1] {
			i++
		} else {
			break
		}
	}

	if i == 1 {
		return false
	}

	for i < n {
		if arr[i] > arr[i-1] {
			dec = true
			i++
		} else {
			return false
		}
	}

	return dec
}

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	first := true
	inc := false
	dec := false
	for i := 1; i < n; {
		if first {
			if arr[i] > arr[i-1] {
				inc = true
				i++
			} else {
				first = false
			}
		} else {
			if arr[i] < arr[i-1] {
				dec = true
				i++
			} else {
				return false
			}
		}
	}

	return dec && inc
}
