package queue_stack

type MovingAverage struct {
	iq    *intqueue
	total float64
	cnt   int
}

func Constructor346(size int) MovingAverage {
	iq := newintqueue(size)
	return MovingAverage{
		iq: iq,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if !this.iq.isfull() {
		this.total = this.total + float64(val)
		this.cnt++
		this.iq.add(val)
		return this.total / float64(this.cnt)
	} else {
		last, _ := this.iq.remove()
		this.total = this.total - float64(last) + float64(val)
		this.iq.add(val)
		return this.total / float64(this.cnt)
	}
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

type intqueue struct {
	arr  []int
	head int
	tail int
	capc int
	size int
}

// head -> [1 , 2 , 3] -> tail
// write at tail
// read at head
func newintqueue(size int) *intqueue {
	return &intqueue{
		arr:  make([]int, size),
		head: 0,
		tail: -1,
		size: 0,
		capc: size,
	}
}

func (q *intqueue) add(val int) bool {
	if q.isfull() {
		return false
	}
	// move tail to next and write
	q.tail = (q.tail + 1) % q.capc
	q.arr[q.tail] = val
	q.size++
	return true
}

func (q *intqueue) remove() (int, bool) {
	if q.isempty() {
		return 0, false
	}

	item := q.arr[q.head]
	q.head = (q.head + 1) % q.capc
	q.size--
	return item, true
}

func (q *intqueue) isfull() bool {
	return q.size == q.capc
}

func (q *intqueue) isempty() bool {
	return q.size == 0
}
