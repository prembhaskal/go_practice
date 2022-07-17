package queue_stack

type MyCircularQueue struct {
	arr  []int
	size int
	head int
	tail int
	k    int
}

func Constructor622(k int) MyCircularQueue {
	return MyCircularQueue{
		arr:  make([]int, k),
		size: 0,
		head: 0,
		tail: -1, // start tail with -1 so that on 1st enqueue, head and tail point to same and continue properly then onwars.
		k:    k,
	}
}

// head 1 - 2 - 3 - 4 - tail
// remove from head
// add at tail

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	// tail points to current, so advance and add.
	this.tail = (this.tail + 1) % this.k
	this.arr[this.tail] = value
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	// zero current position and go to next
	this.arr[this.head] = 0
	this.head = (this.head + 1) % this.k
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.k
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
