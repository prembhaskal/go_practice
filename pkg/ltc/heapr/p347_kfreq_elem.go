package heapr

import "fmt"

// TODO - do using just array, O(n) time complexity
func topKFrequent(nums []int, k int) []int {
	hist := make(map[int]int)
	for _, v := range nums {
		hist[v]++
	}

	numfreqs := make([]numfreq, 0)
	for k, v := range hist {
		numfreqs = append(numfreqs, newnumfreq(k, v))
	}

	// klog(n) complexity / klog(k) complexity
	mh := newheap347(numfreqs)
	res := make([]int, 0)
	for ; k > 0; k-- {
		res = append(res, mh.extractMax().num)
	}
	return res
}

type numfreq struct {
	num  int
	freq int
}

func newnumfreq(num, freq int) numfreq {
	return numfreq{num: num, freq: freq}
}

func (n numfreq) Less(other numfreq) bool {
	if n.freq < other.freq {
		return true
	}
	return false
}

type heap347 struct {
	ar   []numfreq
	size int
}

func newheap347(nums []numfreq) *heap347 {
	m := &heap347{
		ar:   make([]numfreq, len(nums)+1),
		size: len(nums),
	}
	m.ar[0] = numfreq{}
	copy(m.ar[1:], nums)
	m.heapify()
	return m
}

func (m *heap347) parent(i int) int {
	return i / 2
}

// 1 based indices
func (m *heap347) left(i int) int {
	return 2 * i
}

func (m *heap347) right(i int) int {
	return 2*i + 1
}

func (m *heap347) heapify() {
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

func (m *heap347) maxheapify(i int) {
	l := m.left(i)
	r := m.right(i)

	largest := i
	if l <= m.size && !m.ar[l].Less(m.ar[i]) {
		largest = l
	}
	if r <= m.size && !m.ar[r].Less(m.ar[largest]) {
		largest = r
	}
	if largest != i {
		// swap i with largest
		m.ar[i], m.ar[largest] = m.ar[largest], m.ar[i]
		// next check on largest,
		m.maxheapify(largest)
	}
}

func (m *heap347) String() string {
	return fmt.Sprintf("%v", m.ar)
}

func (m *heap347) extractMax() numfreq {
	if m.size <= 0 {
		panic("heap underflow")
	}

	max := m.ar[1]
	m.ar[1] = m.ar[m.size]
	m.size--
	m.maxheapify(1)
	return max
}
