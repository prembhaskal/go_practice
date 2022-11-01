package heapr

import "fmt"

// TODO : try to solve using BST, people are solving with BST it seems

type MedianFinder struct {
	mh  *p295minheap
	xh  *p295maxheap
	med float64
}

func Constructor() MedianFinder {
	return MedianFinder{
		mh: newp295minheap(nil),
		xh: newp295maxheap(nil),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// max_heap -- min heap
	// first add element to max heap,
	// if size(max-heap) > size(min-heap), then extractMax -> insertToMin

	fmt.Printf("before max size: %d, min size: %d, num: %d\n", this.xh.size, this.mh.size, num)
	fmt.Printf("before max heap: %v\n", this.xh.ar[1:this.xh.size+1])
	fmt.Printf("before min heap: %v\n", this.mh.ar[1:this.mh.size+1])

	this.xh.insert(num)

	fmt.Printf("max size: %d, min size: %d, num: %d\n", this.xh.size, this.mh.size, num)
	fmt.Printf("max heap: %v\n", this.xh.ar[1:this.xh.size+1])
	fmt.Printf("min heap: %v\n", this.mh.ar[1:this.mh.size+1])
	if this.mh.size > 0 {
		fmt.Printf("max peek: %d, min peek %d\n", this.xh.peekMax(), this.mh.peekMin())
	}

	max := this.xh.extractMax()
	this.mh.insert(max)

	if this.mh.size > this.xh.size+1 { // if min is too loaded
		min := this.mh.extractMin()
		this.xh.insert(min)
	}

	// if (this.xh.size+this.mh.size)%2 == 0 {
	// 	if this.xh.size > this.mh.size {
	// 		max := this.xh.extractMax()
	// 		this.mh.insert(max)
	// 	}
	// 	this.med = float64(this.xh.peekMax()+this.mh.peekMin()) / 2.0
	// } else {
	// 	if this.xh.size > this.mh.size+1 {
	// 		max := this.xh.extractMax()
	// 		this.mh.insert(max)
	// 	}
	// 	this.med = float64(this.xh.peekMax()) / 1.0
	// }
}

func (this *MedianFinder) FindMedian() float64 {
	if this.mh.size > this.xh.size {
		return float64(this.mh.peekMin()) / 1.0
	}
	return float64(this.xh.peekMax()+this.mh.peekMin()) / 2.0
}

type p295minheap struct {
	ar   []int
	size int
}

func newp295minheap(nums []int) *p295minheap {
	m := &p295minheap{
		ar:   make([]int, len(nums)+1),
		size: len(nums),
	}
	copy(m.ar[1:], nums)
	for i := m.size / 2; i > 0; i-- {
		m.minHeapify(i)
	}
	return m
}

func (m *p295minheap) parent(i int) int {
	return i / 2
}

func (m *p295minheap) left(i int) int {
	return 2 * i
}

func (m *p295minheap) right(i int) int {
	return 2*i + 1
}

func (m *p295minheap) minHeapify(i int) {
	l := m.left(i)
	r := m.right(i)
	minimum := i
	if l <= m.size && m.ar[l] < m.ar[minimum] {
		minimum = l
	}
	if r <= m.size && m.ar[r] < m.ar[minimum] {
		minimum = r
	}
	if minimum != i {
		m.ar[i], m.ar[minimum] = m.ar[minimum], m.ar[i]
		m.minHeapify(minimum)
	}
}

func (m *p295minheap) insert(val int) {
	m.size++
	m.ar = append(m.ar, val)
	// keep swapping up until heap prop restored
	i := m.size
	par := m.parent(i)
	for i > 1 && m.ar[par] > m.ar[i] {
		m.ar[par], m.ar[i] = m.ar[i], m.ar[par]
		i = par
		par = m.parent(i)
	}
}

func (m *p295minheap) peekMin() int {
	return m.ar[1]
}

func (m *p295minheap) extractMin() int {
	min := m.ar[1]
	m.ar[1] = m.ar[m.size]
	m.size--
	m.ar = m.ar[:m.size+1]
	m.minHeapify(1)
	return min
}

type p295maxheap struct {
	ar   []int
	size int
}

func newp295maxheap(nums []int) *p295maxheap {
	x := &p295maxheap{
		ar:   make([]int, len(nums)+1),
		size: len(nums),
	}
	copy(x.ar[1:], nums)
	for i := x.size / 2; i > 1; i-- {
		x.maxHeapify(i)
	}
	return x
}

func (x *p295maxheap) parent(i int) int {
	return i / 2
}

func (x *p295maxheap) left(i int) int {
	return 2 * i
}

func (x *p295maxheap) right(i int) int {
	return 2*i + 1
}

func (x *p295maxheap) maxHeapify(i int) {
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

func (x *p295maxheap) peekMax() int {
	return x.ar[1]
}

func (x *p295maxheap) extractMax() int {
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

func (x *p295maxheap) insert(val int) {
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
