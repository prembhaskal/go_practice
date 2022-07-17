package queue_stack

import (
	"fmt"
)

// https://go.dev/play/p/8ajtaWTZfPO
func main() {
	fmt.Println("Hello, 世界")

	// stack as queue.
	stck := newstk()
	stck.push(1)
	stck.push(2)
	stck.push(3)
	fmt.Printf("stck: %s\n", stck)
	var val int
	val = stck.pop()
	fmt.Printf("pop: %d\n", val)
	fmt.Printf("stck: %s\n", stck)
}

type stk struct {
	arr []int
}

func newstk() *stk {
	return &stk{
		arr: make([]int, 0),
	}
}

func (s *stk) push(val int) {
	s.arr = append(s.arr, val)
}

func (s *stk) pop() int {
	if s.isempty() {
		panic("empty stack")
	}
	n := len(s.arr)
	val := s.arr[n-1]
	s.arr = s.arr[:n-1]
	return val
}

func (s *stk) peek() int {
	return s.arr[len(s.arr)-1]
}

func (s *stk) isempty() bool {
	return len(s.arr) == 0
}

func (s *stk) String() string {
	return fmt.Sprintf("%v", s.arr)
}
