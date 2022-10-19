package heapr

import "fmt"

func findKthLargest(nums []int, k int) int {
	mh := newintmaxheap(nums)
	kth := -1
	for i := 1; i <= k; i++ {
		kth = mh.extractMax()
	}
	return kth
}

type intmaxheap struct {
	ar   []int
	size int
}

func newintmaxheap(nums []int) *intmaxheap {
	m := &intmaxheap{
		ar:   make([]int, len(nums)+1),
		size: len(nums),
	}
	m.ar[0] = -123
	copy(m.ar[1:], nums)
	m.heapify()
	return m
}

func (m *intmaxheap) parent(i int) int {
	return i / 2
}

// 1 based indices
func (m *intmaxheap) left(i int) int {
	return 2 * i
}

func (m *intmaxheap) right(i int) int {
	return 2*i + 1
}

func (m *intmaxheap) heapify() {
	for i := m.size / 2; i > 0; i-- {
		m.maxheapify(i)
	}
}

/*
from CLRS
MAX-HEAPIFY.A; i /
1 l D LEFT.i /
2 r D RIGHT.i /
3 if l <= A:heap-size and A[l]> A[i]
4   largest = l
5 else largest = i
6 if r <= A:heap-size and A[r] > A[largest]
7   largest = r
8 if largest != i
9    exchange A[i] with A[largest]
10   MAX-HEAPIFY(A, largest)

*/

func (m *intmaxheap) maxheapify(i int) {
	l := m.left(i)
	r := m.right(i)

	largest := i
	if l <= m.size && m.ar[l] > m.ar[i] {
		largest = l
	}
	if r <= m.size && m.ar[r] > m.ar[largest] {
		largest = r
	}
	if largest != i {
		// swap i with largest
		m.ar[i], m.ar[largest] = m.ar[largest], m.ar[i]
		// next check on largest,
		m.maxheapify(largest)
	}
}

func (m *intmaxheap) String() string {
	return fmt.Sprintf("%v", m.ar)
}

func (m *intmaxheap) extractMax() int {
	if m.size <= 0 {
		panic("heap underflow")
	}

	max := m.ar[1]
	m.ar[1] = m.ar[m.size]
	m.size--
	m.maxheapify(1)
	return max
}

type intminheap struct {
	ar   []int
	size int
}

func newintminheap(nums []int) *intminheap {
	mh := &intminheap{
		// start with index 1
		ar:   make([]int, len(nums)+1),
		size: len(nums),
	}
	mh.ar[0] = -123
	copy(mh.ar[1:], nums)
	mh.heapify()
	return mh
}

func (m *intminheap) left(i int) int {
	return 2 * i
}

func (m *intminheap) right(i int) int {
	return 2*i + 1
}

func (m *intminheap) parent(i int) int {
	return i / 2
}

func (m *intminheap) heapify() {
	for i := m.size / 2; i > 0; i-- {
		m.minHeapify(i)
	}
}

func (m *intminheap) String() string {
	return fmt.Sprintf("%v", m.ar)
}

// minHeapify is similar to max-heapify from CLRS above, but comparisons are reverse, since min-heap.
func (m *intminheap) minHeapify(i int) {
	l := m.left(i)
	r := m.right(i)

	// find smallest between left and right child, swap current node with that.
	smallest := i
	if l <= m.size && m.ar[l] < m.ar[i] {
		smallest = l
	}
	if r <= m.size && m.ar[r] < m.ar[smallest] {
		smallest = r
	}

	if smallest != i {
		m.ar[i], m.ar[smallest] = m.ar[smallest], m.ar[i]
		// current node is heap now, check next child level where we swapped.
		m.minHeapify(smallest)
	}
}

func (m *intminheap) extractMin() int {
	if m.size < 1 {
		panic("heap underflow")
	}
	min := m.ar[1]
	m.ar[1] = m.ar[m.size]
	m.size--
	m.minHeapify(1)
	return min
}
