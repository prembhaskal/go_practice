package queue_stack

// Explanation - using Stack as Monotonic stack, we have the next greater number or -1
// using map, since given that all nos. in array are unique.
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stk := newstack()
	stk.push(-1)

	nextmax := make(map[int]int)
	for j := len(nums2) - 1; j >= 0; j-- {
		num := nums2[j]
		for {
			top := stk.peek()
			if num < top || top == -1 {
				stk.push(num)
				nextmax[num] = top
				break
			} else {
				stk.pop()
			}
		}
	}

	ans := make([]int, 0)
	for _, v := range nums1 {
		ans = append(ans, nextmax[v])
	}

	return ans
}

type stack struct {
	ar []int
}

func newstack() *stack {
	return &stack{
		ar: make([]int, 0),
	}
}

func (s *stack) push(a int) {
	s.ar = append(s.ar, a)
}

func (s *stack) peek() int {
	if len(s.ar) == 0 {
		panic("empty stack")
	}
	return s.ar[len(s.ar)-1]
}

func (s *stack) pop() int {
	if len(s.ar) == 0 {
		panic("pop on empty stack")
	}
	last := s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1] // remove last element
	return last
}

func (s *stack) size() int {
	return len(s.ar)
}
