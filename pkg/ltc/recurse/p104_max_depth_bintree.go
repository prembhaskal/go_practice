package recurse

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return maxDepthRec(root)
}

func maxDepthRec(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + maxDepthRec(root.Left)
	right := 1 + maxDepthRec(root.Right)

	if left > right {
		return left
	}
	return right
}

func maxDepthIter(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	st := newstack1()
	st.push(newlevelnode(root, 1))
	for !st.isempty() {
		elem := st.pop().(*levelnode)
		left := elem.node.Left
		if left != nil {
			st.push(newlevelnode(left, elem.level+1))
		}
		right := elem.node.Right
		if right != nil {
			st.push(newlevelnode(right, elem.level+1))
		}

		if elem.level > max {
			max = elem.level
		}
	}
	return max
}

type levelnode struct {
	node  *TreeNode
	level int
}

func newlevelnode(node *TreeNode, l int) *levelnode {
	return &levelnode{
		node:  node,
		level: l,
	}
}

type t1 interface{}

type stack1 struct {
	ar []t1
}

func newstack1() *stack1 {
	return &stack1{
		ar: make([]t1, 0),
	}
}

func (s *stack1) push(elem t1) {
	s.ar = append(s.ar, elem)
}

func (s *stack1) pop() t1 {
	var elem t1
	if s.isempty() {
		return elem
	}
	n := len(s.ar)
	elem = s.ar[n-1]
	s.ar = s.ar[:n-1]
	return elem
}

func (s *stack1) isempty() bool {
	return len(s.ar) == 0
}
