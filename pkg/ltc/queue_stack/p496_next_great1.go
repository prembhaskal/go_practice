package queue_stack

import (
	gstack "github.com/emirpasic/gods/stacks/arraystack"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// create monotonic stack for nums2
	stk := gstack.New()
	stk.Push(-1)

	nextBig := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		// check top of stack, if it is bigger than current, then mark that as next big in map
		for {
			top, _ := stk.Peek()
			topn := top.(int)
			if top == -1 || topn > num {
				stk.Push(num)
				nextBig[num] = topn
				break
			} else {
				// keep removing until we get a bigger number
				stk.Pop()
			}
		}
	}
	ans := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = nextBig[nums1[i]]
	}
	return ans
}

// Explanation - using Stack as Monotonic stack, we have the next greater number or -1
// using map, since given that all nos. in array are unique.
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
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
