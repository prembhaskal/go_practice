package heapr

func getSkyline(buildings [][]int) [][]int {


// approach    
    allpts := make([]point, 0)
    // sort points
    for _, build := range buildings {
        allpts = append(allpts, point{build[0], build[2]})
        allpts = append(allpts, point{build[1], build[2]})
    }
    
    pts := &points{allpts}
    
    sort.Sort(pts)
    
    fmt.Printf("all points: %v\n", pts)
    
// sweep line
// for every removal, check if tallest building removed, update the skyline
// take care to remove old taller build, left of the current sweep position, wherever it is seen.
    mh := newp218maxheap(nil)
    
    outline := make([][]int, 0)
    
    for _, v := range pts {
        
    }
    
    return nil    
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
    return p.pts[i].y < p.pts[j].y
}

func (p *points) Swap(i, j int) {
    p.pts[i], p.pts[j] = p.pts[j], p.pts[i]
}

func (p *points) Len() int {
    return len(p.pts)
}

type point struct {
    x int
    y int
}

type p218maxheap struct {
	ar   []int
	size int
}

func newp218maxheap(nums []int) *p218maxheap {
	x := &p218maxheap{
		ar:   make([]int, len(nums)+1),
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
	if l <= x.size && x.ar[l] > x.ar[max] {
		max = l
	}
	if r <= x.size && x.ar[r] > x.ar[max] {
		max = r
	}
	if max != i {
		// swap
		x.ar[max], x.ar[i] = x.ar[i], x.ar[max]
		x.maxHeapify(max)
	}
}

func (x *p218maxheap) peekMax() int {
	return x.ar[1]
}

func (x *p218maxheap) extractMax() int {
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

func (x *p218maxheap) insert(val int) {
	x.ar = append(x.ar, val)
	x.size++

	i := x.size
	par := x.parent(i)
	for i > 1 && x.ar[par] < x.ar[i] {
		x.ar[par], x.ar[i] = x.ar[i], x.ar[par]
		i = par
		par = x.parent(i)
	}
}