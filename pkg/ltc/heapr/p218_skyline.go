package heapr

import (
	"fmt"
	"sort"
)

//  TODO - accepted, but optimize here.
func getSkyline(buildings [][]int) [][]int {
	// approach
	allpts := make([]point, 0)
	// sort points
	for _, build := range buildings {
		allpts = append(allpts, point{build[0], build[2], build[1]})
		allpts = append(allpts, point{build[1], -build[2], build[1]}) // add negative height to indicate this is falling edge.
	}

	pts := &points{allpts}

	sort.Sort(pts)

	// fmt.Printf("all points: %v\n", pts)

	// sweep line
	// for every removal, check if tallest building removed, update the skyline
	// take care to remove old taller build, left of the current sweep position, wherever it is seen.
	mh := newp218maxheap(nil)

	outline := make([][]int, 0)

	for _, v := range pts.pts {
		fmt.Printf("building to consider, x: %d, height: %d, right: %d\n", v.x, v.y, v.e)

		fmt.Printf("heap content: %v\n", mh.ar)

		if v.y > 0 { // build start
			// if empty heap - first building, also mark
			if mh.size == 0 {
				outline = append(outline, []int{v.x, v.y})
				mh.insert(v)
				continue
			}
			prevTall := mh.peekMax()
			// new building taller than previous tallest, start new outline, also mark the current building
			if v.y > prevTall.y {
				outline = append(outline, []int{v.x, v.y})
				mh.insert(v)
				continue
			} else if v.y == prevTall.y { // new building same are previous tallest, outline no change, but mark current building
				mh.insert(v)
				continue
			} else if v.y < prevTall.y { // new building smaller than previous tallest, outline no change , mark currnt building
				mh.insert(v)
				continue
			}
		} else { // building end.
			if mh.size > 0 {
				prevTall := mh.peekMax()
				// if curr building same as prev. tallest
				if -v.y == prevTall.y {
					// clean up older buildings, which are no longer valid
					for mh.size > 0 && mh.peekMax().e <= v.x {
						mh.extractMax()
					}

					// if no more building, then mark outline 0
					if mh.size == 0 {
						outline = append(outline, []int{v.x, 0})
						continue
					}

					// still some buildings left, these are definitely valid. add the outline if not same height
					if mh.size > 0 {
						nexttall := mh.peekMax()
						if nexttall.y != -v.y { // if current falling edge, is same as another building still live.
							outline = append(outline, []int{v.x, nexttall.y})
							continue
						}

						// if there are building same height as current, still ongoing, just ignore them, outline continues
					}
				} else {
					// current ending builing is not prev. tallest, because some taller building already started
					// nothing we can do here. old invalid will be cleared later on.
				}
			}
		}
	}

	return outline
}

type points struct {
	pts []point
}

func (p *points) Less(i, j int) bool {
	if p.pts[i].x < p.pts[j].x {
		return true
	}
	if p.pts[i].x > p.pts[j].x {
		return false
	}
	return p.pts[i].y > p.pts[j].y // for building falling, starting at same point, process starting building first.
}

func (p *points) Swap(i, j int) {
	p.pts[i], p.pts[j] = p.pts[j], p.pts[i]
}

func (p *points) Len() int {
	return len(p.pts)
}

type point struct { // rename to building make more sense??
	x int
	y int
	e int // right edge of build
}

type p218maxheap struct {
	ar   []point
	size int
}

func newp218maxheap(nums []point) *p218maxheap {
	x := &p218maxheap{
		ar:   make([]point, len(nums)+1),
		size: len(nums),
	}
	copy(x.ar[1:], nums)
	for i := x.size / 2; i > 1; i-- {
		x.maxHeapify(i)
	}
	return x
}

func (x *p218maxheap) parent(i int) int {
	return i / 2
}

func (x *p218maxheap) left(i int) int {
	return 2 * i
}

func (x *p218maxheap) right(i int) int {
	return 2*i + 1
}

func (x *p218maxheap) maxHeapify(i int) {
	l := x.left(i)
	r := x.right(i)
	max := i
	if l <= x.size && x.ar[l].y > x.ar[max].y {
		max = l
	}
	if r <= x.size && x.ar[r].y > x.ar[max].y {
		max = r
	}
	if max != i {
		// swap
		x.ar[max], x.ar[i] = x.ar[i], x.ar[max]
		x.maxHeapify(max)
	}
}

func (x *p218maxheap) peekMax() point {
	return x.ar[1]
}

func (x *p218maxheap) extractMax() point {
	if x.size < 1 {
		panic("underflow max heap")
	}
	max := x.ar[1]
	x.ar[1] = x.ar[x.size]
	x.size--
	x.ar = x.ar[:x.size+1]
	x.maxHeapify(1)
	return max
}

func (x *p218maxheap) insert(val point) {
	x.ar = append(x.ar, val)
	x.size++

	i := x.size
	par := x.parent(i)
	for i > 1 && x.ar[par].y < x.ar[i].y {
		x.ar[par], x.ar[i] = x.ar[i], x.ar[par]
		i = par
		par = x.parent(i)
	}
}
