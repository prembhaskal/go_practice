package binsearch

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s1 := 0
	s2 := 0
	e1 := len(nums1) - 1
	e2 := len(nums2) - 1
	n1 := e1 - s1 + 1
	n2 := e2 - s2 + 1
	ar1 := nums1[s1 : e1+1]
	ar2 := nums2[s2 : e2+1]

	for n1 >= 2 && n2 >= 2 {
		med1 := findMedian(ar1)
		med2 := findMedian(ar2)
		m11, m12 := findMedianIndex(ar1)
		m21, m22 := findMedianIndex(ar2)

		if med1 == med2 {
			return med1 // or med2 both are same
		} else if med1 < med2 {
			// discard left of med1 and right of med2
			s1 = m11
			e2 = m22
		} else { // med1 > med2
			// discard left of med2 and right of med1
			s2 = m21
			e1 = m12
		}
		n1 = e1 - s1 + 1
		n2 = e2 - s2 + 1

		ar1 = nums1[s1 : e1+1]
		ar2 = nums2[s2 : e2+1]
	}
	if n1 == 0 {
		return findMedian(ar2)
	}
	if n2 == 0 {
		return findMedian(ar1)
	}
	if n1 == 1 && n2 == 1 {
		return avg(ar1[0], ar2[0])
	}
	if n1 == 1 && n2 == 2 {
		return getMedianOneTwo(ar1[0], ar2)
	}
	if n2 == 1 && n1 == 2 {
		return getMedianOneTwo(ar2[0], ar1)
	}
	if n1 == 2 && n2 == 2 {
		return getMedianTwo(ar1, ar2)
	}
	if n1 == 1 && n2 > 2 {
		return getMedianOneMore(ar1[0], ar2)
	}
	if n2 == 1 && n1 > 2 {
		return getMedianOneMore(ar2[0], ar1)
	}

	panic(fmt.Sprintf("this should not happen, n1:%d, n2:%d", n1, n2))
}

func getMedianOneTwo(num int, ar2 []int) float64 {
	if num >= ar2[1] {
		return float64(ar2[1])
	} else if num <= ar2[0] {
		return float64(ar2[0])
	}
	return float64(num)
}

/*
ar1: 7
ar2: 1, 4, 6, 8, 9

net len: even
m_of(7, 4, 6, 8)

ar1: 7 | 9 | 3
ar2: 1, 4, 6, 8, 9, 10

net len: odd
m_of(7, 6, 8)
*/
func getMedianOneMore(num int, arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		return medianOfThree(num, arr[n/2], arr[(n-1)/2])
	} else {
		return medianOfFour(num, arr[n/2-1], arr[n/2], arr[n/2+1])
	}
}

func medianOfThree(a, b, c int) float64 {
	mx := max(max(a, b), c)
	mn := min(min(a, b), c)
	return float64(a + b + c - mx - mn)
}

func medianOfFour(a, b, c, d int) float64 {
	mx := max(max(a, b), max(c, d))
	mn := min(min(a, b), min(c, d))
	return float64(a+b+c+d-mx-mn) / 2.0
}

func getMedianTwo(ar1, ar2 []int) float64 {
	return avg(max(ar1[0], ar2[0]), min(ar1[1], ar2[1]))
}

func avg(a, b int) float64 {
	return float64(a+b) / 2.0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMedian(ar []int) float64 {
	n := len(ar)
	if n%2 == 0 {
		return float64(ar[n/2]+ar[(n-1)/2]) / 2.0
	}
	return float64(ar[n/2])
}

func findMedianIndex(ar []int) (int, int) {
	n := len(ar)
	if n%2 == 0 {
		return (n - 1) / 2, n / 2
	}
	return n / 2, n / 2
}
