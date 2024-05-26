package slidewin

// solved using a deque
// deque is kept at max size of 'k' by invariant in the loop
// add new items to last of deque, while adding we remove any existing smaller items we encounter from LAST
// smaller items don't matter as bigger item dominates the answer.
// also this keep the queue in sorted order from LAST to FIRST (smallest item in LAST, biggest item at FIRST)
// this is very similar to monotonic stack concept.
// read from FIRST to get the answer at current index,
// while reading from FIRST, we discard any items not part of current sliding window
// TODO Use simple slice instead of deque, store item index instead of value in it.
func maxSlidingWindow(nums []int, k int) []int {

	deq := NewDeq()

	// add to last
	// if new item being pushed is bigger than existing, remove existing (repeat)
	for i := 0; i < k; i++ {
		num := nums[i]
		for deq.size > 0 && num > deq.peekLast().(pair).val {
			deq.removeLast()
		}
		deq.addLast(pair{num, i})
	}

	ans := make([]int, 0)
	ans = append(ans, deq.peekFirst().(pair).val)

	for i := k; i < len(nums); i++ {
		num := nums[i]
		for deq.size > 0 && num > deq.peekLast().(pair).val {
			deq.removeLast()
		}
		deq.addLast(pair{num, i})

		// clean up items from first if out of index
		// add peek to ans
		for {
			toppair := deq.peekFirst().(pair)
			if toppair.idx > i-k {
				break
			}
			deq.removeFirst()
		}

		ans = append(ans, deq.peekFirst().(pair).val)
	}

	return ans
}

type pair struct {
	val int
	idx int
}

// (first) HEAD -> item1 -> item2 -> item3 -> TAIL (last)
type deque struct {
	ar   []interface{}
	size int
	head *dnode
	tail *dnode
}

var HEAD = &dnode{
	prev: nil,
}

var TAIL = &dnode{
	next: nil,
}

type dnode struct {
	val  interface{}
	next *dnode
	prev *dnode
}

func NewDeq() *deque {
	HEAD.next = TAIL
	TAIL.prev = HEAD
	deq := &deque{
		ar:   make([]interface{}, 0),
		size: 0,
		head: HEAD,
		tail: TAIL,
	}
	return deq
}

// add first
func (d *deque) addFirst(item any) {
	newnode := &dnode{
		val:  item,
		prev: HEAD,
		next: HEAD.next,
	}

	nextitem := HEAD.next
	HEAD.next = newnode
	nextitem.prev = newnode
	d.size++
}

// remove first
func (d *deque) removeFirst() interface{} {
	if d.size == 0 {
		panic("empty queue")
	}
	first := HEAD.next

	second := first.next
	HEAD.next = second
	second.prev = HEAD

	d.size--

	return first.val
}

// peek first
func (d *deque) peekFirst() any {
	d.sizecheck()

	return HEAD.next.val
}

// add last
func (d *deque) addLast(item any) {
	lastnode := &dnode{
		val:  item,
		next: TAIL,
		prev: TAIL.prev,
	}

	prevnode := TAIL.prev
	prevnode.next = lastnode

	TAIL.prev = lastnode

	d.size++
}

// remove last
func (d *deque) removeLast() any {
	d.sizecheck()

	last := TAIL.prev
	secondlast := last.prev

	secondlast.next = TAIL
	TAIL.prev = secondlast

	d.size--
	return last.val
}

// peek last
func (d *deque) peekLast() any {
	d.sizecheck()

	return TAIL.prev.val
}

func (d *deque) sizecheck() {
	if d.size == 0 {
		panic("empty queue")
	}
}
