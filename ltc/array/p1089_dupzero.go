package array

// https://leetcode.com/problems/duplicate-zeros/

// Explanation
// [1, 0, 0, 0, 2, 3]


// cnt = 3

// // 3 elements to be removed from back
// [1, 0, 0, 0, 0, 0, 0, 2 , 3]  // eg. aisa dikhta actual array, if it could be extended

// so 3 is removed   					i = 5
//    2 is removed   					i = 4
//    0 is removed  but only once      i = 3 // note that countonce using a flag.

// so we have one pointer at the last after removal at i == 3

// one pointer at i // from above calc.
// one pointer at j // last index of array

// keep iterating from back,
// 	if 0, fill twice, reduce j twice
// 	if 1, fill once, reduce j once
	
// 	reduce i
func duplicateZeros(arr []int) {
	// find no. of zeroes within array.
	cnt := -1
	lastNonZero := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if cnt < 0 && arr[i] != 0 { // mark when we see 1st non zero number.
			cnt = 0 // 1st non zero
			lastNonZero = i
			continue
		}

		if arr[i] == 0 && cnt >= 0 {
			cnt++
		}
	}

	// cnt is total zeros inside, that much space is needed in
	// find which element to being writing from end
	onezero := false
	i := len(arr) - 1
	for ; cnt > 0; i-- {
		if arr[i] != 0 || i > lastNonZero {
			cnt--
		} else {
			// reduce twice for zeros
			cnt = cnt - 2
			if cnt < 0 {
				onezero = true
				break
			}
		}
	}

	// fmt.Printf("value of i after counting: %d\n", i)

	for j := len(arr) - 1; j >= 0; {
		if onezero {
			onezero = false
			arr[j] = 0
			j--
			i--
			continue
		}

		if arr[i] == 0 {
			arr[j] = 0
			j--
			if i < lastNonZero && j >= 0 {
				arr[j] = 0
				j--
			}

		} else {
			arr[j] = arr[i]
			j--
		}
		i--

	}
}

func duplicateZeros1(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			insertZero(i, arr)
			i++
		}
	}
}

func insertZero(idx int, arr []int) {
	for i := len(arr) - 1; i > idx; i-- {
		arr[i] = arr[i-1]
	}
	arr[idx] = 0
}
