package wc_278

// import "fmt"

//    [0, 0, 1, 0, _]

// z  [0, 1, 2, 2, 3]

// on [1, 1, 1, 0, 0]

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	zerocnt := make([]int, n+1)

	for i := 0; i < n; i++ {
		zerocnt[i+1] = zerocnt[i]
		if nums[i] == 0 {
			zerocnt[i+1] = zerocnt[i+1] + 1
		}
	}

	onecnt := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		onecnt[i] = onecnt[i+1]
		if nums[i] == 1 {
			onecnt[i] = onecnt[i] + 1
		}
	}

	max := -1
	for i := 0; i < n+1; i++ {
		tmp := zerocnt[i] + onecnt[i]
		if tmp > max {
			max = tmp
		}
	}

	// fmt.Printf("zerocnt: %v\n", zerocnt)
	// fmt.Printf("onecnt: %v\n", onecnt)
	// fmt.Printf("max : %d\n", max)

	arr := make([]int, 0)
	for i := 0; i < n+1; i++ {
		if max == (zerocnt[i] + onecnt[i]) {
			arr = append(arr, i)
		}
	}

	return arr
}
