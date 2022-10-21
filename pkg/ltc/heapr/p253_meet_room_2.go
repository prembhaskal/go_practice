package heapr

import "sort"

func minMeetingRooms(intervals [][]int) int {
	intvs := make([]interval, 0)
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		intvs = append(intvs, interval{start, 1}) // add room for start,
		intvs = append(intvs, interval{end, -1})  // free room for end.
	}

	data := &sortintvs{
		ar: intvs,
	}

	sort.Sort(data)

	// fmt.Printf("sorted data: %v\n", data.ar)

	max := 0
	rooms := 0
	for i := 0; i < len(data.ar); i++ {
		curr := data.ar[i]
		// check current time, if it is starting, add a room, if it is ending, free a room.
		rooms = rooms + curr.cnt
		if rooms > max {
			max = rooms
		}
	}

	return max
}

type sortintvs struct {
	ar []interval
}

func (s *sortintvs) Less(i, j int) bool {
	if s.ar[i].num < s.ar[j].num {
		return true
	} else if s.ar[i].num > s.ar[j].num {
		return false
	}

	// for same interval, keep ending slots before starting slots.
	return s.ar[i].cnt < s.ar[j].cnt
	// return false
}

func (s *sortintvs) Len() int {
	return len(s.ar)
}

func (s *sortintvs) Swap(i, j int) {
	s.ar[i], s.ar[j] = s.ar[j], s.ar[i]
}

type interval struct {
	num int
	cnt int
}
