package recurse

import (
	"testing"
)

func TestCourses1(t *testing.T) {
	n := 4
	relations := [][]int{{2,1},{3,1},{1,4}}
	k := 2
	var ans int
	
	ans = minNumberOfSemesters(n, relations, k )
	t.Errorf("answer is %d\n", ans)
}

func TestCourses2(t *testing.T) {
	n := 5
	relations := [][]int{{2,1},{3,1},{4,1}, {1,5}}
	k := 2
	var ans int
	
	ans = minNumberOfSemesters(n, relations, k )
	t.Errorf("answer is %d\n", ans)
}

func TestCourses3(t *testing.T) {
	n := 5
	relations := [][]int{}
	k := 2
	var ans int
	
	ans = minNumberOfSemesters(n, relations, k )
	t.Errorf("answer is %d\n", ans)
}
func TestCourse4(t *testing.T) {
	n := 13
	relations := [][]int{{12,8},{2,4},{3,7},{6,8},{11,8},{9,4},{9,7},{12,4},{11,4},{6,4},{1,4},{10,7},{10,4},{1,7},{1,8},{2,7},{8,4},{10,8},{12,7},{5,4},{3,4},{11,7},{7,4},{13,4},{9,8},{13,8}}
	k := 9
	var ans int
	
	ans = minNumberOfSemesters(n, relations, k )
	t.Errorf("answer is %d\n", ans)
}

func TestGenMask(t *testing.T) {
	n := 13
	totalmasks := genmask(n, 0, 1)
	t.Errorf("total masks are %d\n", totalmasks)
}