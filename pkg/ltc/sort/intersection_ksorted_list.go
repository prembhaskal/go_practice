package sort

import (
	"fmt"
)

func main() {
	fmt.Println("program start")
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{1, 3, 4, 5, 6, 7}
	c := []int{1, 2, 3, 4, 5, 6, 7}
	d := []int{1, 2, 3, 4, 5, 7}
	e := []int{1, 2, 3, 4, 5, 6, 7}
	f := []int{1, 2, 3, 4, 5}
	g := []int{2, 3}

	allNos := [][]int{a,b,c,d,e, f,g}
	inter := intersection(allNos)
	fmt.Printf("intersection is: %v\n", inter)

	fmt.Println("program done")
}

func intersection(allNos [][]int) []int {
	return intersec(allNos, 0, len(allNos)-1)
}

func intersec(allNos [][]int, start, end int) []int {
	if start > end {
		return []int{}
	}
	if start == end {
		return allNos[start]
	}
	mid := (start + end) / 2

	first := intersec(allNos, start, mid)
	second := intersec(allNos, mid+1, end)
	return merge(first, second)
}

func merge(first, second []int) []int {
	i := 0
	j := 0
	ans := make([]int, 0)
	for i < len(first) && j < len(second) {
		if first[i] == second[j] {
			ans = append(ans, first[i])
			i++
			j++
		} else if first[i] < second[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}
