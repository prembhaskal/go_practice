package heapr

import (
	"fmt"
	"sort"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// return furthestMinHeap(heights, bricks, ladders)
	return furthestBinSearch(heights, bricks, ladders)
}

// check if we are able to reach building n/2,
// if yes, search from n/2+1, end
// if no, search from 0 to n/2
func furthestBinSearch(ht []int, br, ld int) int {
	left := 0
	right := len(ht) - 1
	for left < right {
		mid := left + (right-left+1)/2
		if isBuildingReachable(ht, mid, br, ld) {
			// fmt.Printf("building: %d, reachable\n", mid)
			left = mid
		} else {
			// fmt.Printf("building: %d, not reachable\n", mid)
			right = mid - 1
		}
	}

	return left
}

func isBuildingReachable(ht []int, num, br, ld int) bool {
	climbs := make([]int, 0)

	// TODO: optimize, calculate this only once.
	for curr := 0; curr < num; curr++ {
		climb := ht[curr+1] - ht[curr]
		if climb > 0 {
			climbs = append(climbs, climb)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(climbs)))

	// fmt.Printf("for building: %d, climbs: %d\n", num, climbs)

	for i := 0; i < len(climbs); i++ {
		if ld > 0 {
			ld--
			continue
		}
		br = br - climbs[i]
		// fmt.Printf("for building: %d, remaining bricks: %d\n", num, br)
		if br < 0 {
			return false
		}
	}

	return true
}

// heap based approach
// strategy: eventual goal is to assign the max climbs to ladder, remaining to bricks
// first assign all climbs to ladder
// for each, climb, if climb > minimum assigned ladder, reassign ladder to this.
//                  else assign bricks
//       stop and return current building, if we run out of bricks/ladder.

func furthestMinHeap(heights []int, bricks, ladders int) int {
	lh := newp1642heap(nil)
	curr := 0
out:
	for ; curr < len(heights)-1; curr++ {
		// check if climb
		// next := curr + 1
		climb := heights[curr+1] - heights[curr]
		if climb <= 0 {
			continue
		}

		if ladders > 0 {
			lh.minHeapInsert(climb)
			ladders--
			continue
		}

		// no ladders, check if we can swap with ladder assignment
		if lh.Size() > 0 && climb > lh.peekMin() {
			minLadder := lh.peekMin()
			lh.updateTop(climb)
			climb = minLadder
		}

		// assign bricks if we have reached here
		if bricks < climb {
			break out
		}

		bricks = bricks - climb
	}

	return curr
}

type p1642heap struct {
	ar   []int
	size int
}

func newp1642heap(nums []int) *p1642heap {
	mh := &p1642heap{
		ar:   make([]int, len(nums)+1),
		size: len(nums),
	}

	mh.ar[0] = 11111
	copy(mh.ar[1:], nums)

	for i := mh.size / 2; i > 0; i-- {
		mh.minHeapify(i)
	}

	return mh
}

func (h *p1642heap) Size() int {
	return h.size
}

func (h *p1642heap) String() string {
	return fmt.Sprintf("%v", h.ar)
}

func (h *p1642heap) parent(idx int) int {
	return idx / 2
}

func (h *p1642heap) left(idx int) int {
	return 2 * idx
}

func (h *p1642heap) right(idx int) int {
	return 2*idx + 1
}

func (h *p1642heap) minHeapify(idx int) {
	l := h.left(idx)
	r := h.right(idx)
	smallest := idx
	if l <= h.size && h.ar[l] < h.ar[smallest] {
		smallest = l
	}
	if r <= h.size && h.ar[r] < h.ar[smallest] {
		smallest = r
	}
	if smallest != idx {
		// swap and continue with next
		h.ar[smallest], h.ar[idx] = h.ar[idx], h.ar[smallest]
		h.minHeapify(smallest)
	}
}

func (h *p1642heap) minHeapInsert(key int) {
	h.ar = append(h.ar, key)
	h.size++
	h.ar[h.size] = 1 << 31 // 2^31 as infinity.
	h.decreaseKey(h.size, key)
}

func (h *p1642heap) peekMin() int {
	if h.size == 0 {
		panic("nothing to peek")
	}
	return h.ar[1]
}

func (h *p1642heap) extractMin() int {
	min := h.ar[1]
	h.ar[1] = h.ar[h.size]
	h.ar[h.size] = min
	h.size--
	h.minHeapify(1)
	return min
}

func (h *p1642heap) updateTop(newmin int) {
	h.ar[1] = newmin
	h.minHeapify(1)
}

func (h *p1642heap) decreaseKey(idx, key int) {
	if h.ar[idx] < key {
		panic("only decrease supported here")
	}
	h.ar[idx] = key
	par := h.parent(idx)
	for idx > 1 && h.ar[idx] < h.ar[par] {
		// swap with parent and repeat on parent
		h.ar[idx], h.ar[par] = h.ar[par], h.ar[idx]
		idx = par
		par = h.parent(idx)
	}
}
