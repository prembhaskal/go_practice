package binsearch

func findClosestElements(arr []int, k int, x int) []int {
	start := 0
	end := len(arr) - 1
	key := x
	for start+1 < end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			start = mid
			break
		} else if arr[mid] > key {
			end = mid
		} else if arr[mid] < key {
			start = mid
		}
	}

	// start will have element == key , or element just small than key

	// fmt.Printf("closer element is at arr[%d]=%d\n", start, arr[start])

	// close := make([]int, 0)
	left := start
	right := start + 1
	for i := 0; i < k; i++ {
		// var closer int
		if left < 0 { // take right
			// closer = arr[right]
			// close = append(close, closer)
			right++
		} else if right >= len(arr) { // take left
			// closer = arr[left]
			// close = prepend(close, closer)
			left--
		} else if (key - arr[left]) <= (arr[right] - key) {
			// take left
			// closer = arr[left]
			// close = prepend(close, closer)
			left--
		} else { // take right
			// closer = arr[right]
			// close = append(close, closer)
			right++
		}
		// fmt.Printf("add %d to list, i: %d\n", closer, i)
		// fmt.Printf("left: %d, right: %d\n", left, right)
	}

	return arr[left+1 : right]
}

// func abs(a int) int{
//     if a < 0 {
//         return -a
//     }
//     return a
// }

// func prepend(ar []int, x int) []int {
//     ar = append(ar, 0)
//     copy(ar[1:], ar)
//     ar[0] = x
//     return ar
// }
