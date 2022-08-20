package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	return preorderTraversalIter(root)
}

func preorderTraversalRec(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := make([]int, 0)
	nodes = append(nodes, root.Val)

	left := preorderTraversal(root.Left)
	nodes = append(nodes, left...)

	right := preorderTraversal(root.Right)
	nodes = append(nodes, right...)
	return nodes
}

func preorderTraversalIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nodes := make([]int, 0)
	stk := newpreorderstack()
	stk.push(root)

	for !stk.isempty() {
		rt := stk.pop()
		nodes = append(nodes, rt.Val)

		if rt.Right != nil {
			stk.push(rt.Right)
		}
		if rt.Left != nil {
			stk.push(rt.Left)
		}
	}
	return nodes
}

type preorderstack struct {
	ar []*TreeNode
}

func newpreorderstack() *preorderstack {
	return &preorderstack{
		ar: make([]*TreeNode, 0),
	}
}

func (s *preorderstack) push(node *TreeNode) {
	s.ar = append(s.ar, node)
}

func (s *preorderstack) pop() *TreeNode {
	if s.isempty() {
		return nil
	}
	node := s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return node
}

func (s *preorderstack) isempty() bool {
	return len(s.ar) == 0
}
