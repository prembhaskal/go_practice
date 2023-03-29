package queue_stack

//	to check 75
//
// 72
// 74
// 78
// monotonic stack approach - items in stack are sorted.
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stk := newdiststack()
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		cnode := distnode{val: temperatures[i], idx: i}
		for !stk.isempty() {
			tnode, _ := stk.peek()
			if tnode.val >= cnode.val {
				break
			}
			stk.pop()
			dist[tnode.idx] = i - tnode.idx
		}
		stk.push(cnode)
	}

	return dist
}

type distnode struct {
	val int
	idx int
}

type diststack struct {
	ar []distnode
}

func newdiststack() *diststack {
	return &diststack{
		ar: make([]distnode, 0),
	}
}

func (s *diststack) push(n distnode) {
	s.ar = append(s.ar, n)
}

func (s *diststack) peek() (distnode, bool) {
	var dnode distnode
	if s.isempty() {
		return dnode, false
	}
	dnode = s.ar[len(s.ar)-1]
	return dnode, true
}

func (s *diststack) pop() (distnode, bool) {
	var dnode distnode
	if s.isempty() {
		return dnode, false
	}
	dnode = s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return dnode, true
}

func (s *diststack) isempty() bool {
	return len(s.ar) == 0
}
