package queue_stack

type MinStack struct {
	stk  *stk155
	mstk *stk155
}

func Constructor() MinStack {
	return MinStack{
		stk:  newstk155(),
		mstk: newstk155(),
	}
}

func (this *MinStack) Push(val int) {
	this.stk.push(val)
	if this.mstk.isempty() {
		this.mstk.push(val)
		return
	}
	cmin := this.mstk.peek()
	if val < cmin {
		cmin = val
	}
	this.mstk.push(cmin)
}

func (this *MinStack) Pop() {
	this.stk.pop()
	this.mstk.pop()
}

func (this *MinStack) Top() int {
	return this.stk.peek()
}

func (this *MinStack) GetMin() int {
	return this.mstk.peek()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type stk155 struct {
	arr []int
}

func newstk155() *stk155 {
	return &stk155{
		arr: make([]int, 0),
	}
}

func (s *stk155) push(val int) {
	s.arr = append(s.arr, val)
}

func (s *stk155) pop() int {
	n := len(s.arr)
	val := s.arr[n-1]
	s.arr = s.arr[:n-1]
	return val
}

func (s *stk155) peek() int {
	return s.arr[len(s.arr)-1]
}

func (s *stk155) isempty() bool {
	return len(s.arr) == 0
}
