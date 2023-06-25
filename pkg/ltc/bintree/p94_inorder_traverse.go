package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// in order --> in(LEFTSIDE) - ROOT - in(RIGHTSIDE)
func inorderTraversal(root *TreeNode) []int {
	// return inorderTraversalRec(root)
	return inorderTraversalIter(root)
}

func inorderTraversalRec(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := inorderTraversalRec(root.Left)
	arr = append(arr, root.Val)
	arr1 := inorderTraversalRec(root.Right)
	arr = append(arr, arr1...)
	return arr
}

func inorderTraversalIter(root *TreeNode) []int {
	stk := newinorderstack()
	// stk.push(root)
	nodes := make([]int, 0)

	curr := root

	for !stk.isempty() || curr != nil {
		// curr is top, push it and all its left
		for curr != nil {
			stk.push(curr)
			curr = curr.Left
		}

		curr = stk.pop() // got a curr whose left side is all handled by now.
		nodes = append(nodes, curr.Val)

		curr = curr.Right // don't add right yet, make it root and repeat.
	}

	return nodes
}

type inorderstack struct {
	ar []*TreeNode
}

func newinorderstack() *inorderstack {
	return &inorderstack{
		ar: make([]*TreeNode, 0),
	}
}

func (s *inorderstack) push(node *TreeNode) {
	s.ar = append(s.ar, node)
}

func (s *inorderstack) pop() *TreeNode {
	if s.isempty() {
		return nil
	}
	node := s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return node
}

func (s *inorderstack) isempty() bool {
	return len(s.ar) == 0
}
