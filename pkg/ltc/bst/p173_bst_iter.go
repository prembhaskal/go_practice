package bst

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	curr *TreeNode
	stk  *bstStk
}

func Constructor(root *TreeNode) BSTIterator {
	stk := newBstStack()
	// if root != nil {
	//     stk.push(root)
	// }
	return BSTIterator{
		curr: root,
		stk:  stk,
	}
}

func (this *BSTIterator) Next() int {
	currVal := -1
	if this.curr != nil {
		currVal = this.curr.Val
	}
	fmt.Printf("next call, is stack empty: %t, curr val: %d\n", this.stk.isEmpty(), currVal)
	for this.curr != nil {
		this.stk.push(this.curr)
		this.curr = this.curr.Left
	}

	newcurr := this.stk.pop()
	this.curr = newcurr.Right
	fmt.Printf("next call, newcurr: %d\n", newcurr.Val)
	return newcurr.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.curr != nil || !this.stk.isEmpty()
}

type bstStk struct {
	ar []*TreeNode
}

func newBstStack() *bstStk {
	return &bstStk{
		ar: make([]*TreeNode, 0),
	}
}

func (s *bstStk) push(node *TreeNode) {
	s.ar = append(s.ar, node)
}

func (s *bstStk) pop() *TreeNode {
	if s.isEmpty() {
		panic("pop on empty stack.")
	}

	n := len(s.ar)
	popi := s.ar[n-1]
	s.ar = s.ar[:n-1]
	return popi
}

func (s *bstStk) isEmpty() bool {
	return len(s.ar) == 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
