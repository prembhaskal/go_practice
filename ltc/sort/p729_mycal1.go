package sort

import (
	"sort"
)

type MyCalendar struct {
	curr slots
}

func Constructor() MyCalendar {
	return MyCalendar{
		curr: make([]*slot, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	var left *slot
	var right *slot
	// search in curr slots
	for i := 0; i < len(this.curr); i++ {
		slot := this.curr[i]
		if slot.start <= start {
			left = slot
		}
		if right == nil && slot.start > start {
			right = slot
		}
	}

	if left != nil {
		if left.start == start || left.end > start {
			return false
		}
	}

	if right != nil {
		if right.start < end {
			return false
		}
	}

	newslot := &slot{start, end}
	this.curr = append(this.curr, newslot)
	sort.Sort(&this.curr)
	return true
}

type slot struct {
	start int
	end   int
}

type slots []*slot

func (s *slots) Len() int {
	sd := *s
	return len(sd)
}

func (s *slots) Less(i, j int) bool {
	sd := *s
	return sd[i].start < sd[j].start
}

func (s *slots) Swap(i, j int) {
	sd := *s
	sd[i], sd[j] = sd[j], sd[i]
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
