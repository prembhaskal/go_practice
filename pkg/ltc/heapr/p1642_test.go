package heapr

import "testing"

func TestMinHeap1642(t *testing.T) {
	nums := make([]int, 0)
	for i := 30; i > 10; i-- {
		nums = append(nums, i)
	}
	mh := newp1642heap(nums)
	t.Logf("min heap is %s\n", mh)

	mh.minHeapInsert(5)
	t.Logf("min heap after insert is %s\n", mh)

	mh.minHeapInsert(6)
	t.Logf("min heap after insert is %s\n", mh)

	mh.minHeapInsert(31)
	t.Logf("min heap after insert is %s\n", mh)

	for mh.size > 0 {
		min := mh.extractMin()
		t.Logf("min: %d\n", min)
		t.Logf("min heap after extract min %s\n", mh)
	}

	mh1 := newp1642heap(nil)
	t.Logf("empty heap is %s", mh1)
}

func TestFurthestBuilding(t *testing.T) {
	ht := []int{4, 2, 7, 6, 9, 14, 12}
	res := furthestBuilding(ht, 5, 1)
	t.Logf("results is %d\n", res)
}
