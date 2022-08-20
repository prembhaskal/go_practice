package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// visit left -> right -> root
func postorderTraversal(root *TreeNode) []int {
	// return postorderTraversalRec(root)
	return postorderTraversalIter(root)
}

func postorderTraversalRec(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nodes := make([]int, 0)
	left := postorderTraversal(root.Left)
	nodes = append(nodes, left...)

	right := postorderTraversal(root.Right)
	nodes = append(nodes, right...)

	nodes = append(nodes, root.Val)

	return nodes
}

func postorderTraversalIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nodes := make([]int, 0)

	stk := newpostorderstack()
	stk.push(newPostNode(root, false))

	for !stk.isempty() {
		curr := stk.pop()
		if curr.node == nil {
			continue
		}
		if curr.done {
			nodes = append(nodes, curr.node.Val)
		} else {
			stk.push(newPostNode(curr.node, true))
			stk.push(newPostNode(curr.node.Right, false))
			stk.push(newPostNode(curr.node.Left, false))
		}
	}

	return nodes
}

type PostNode struct {
	node *TreeNode
	done bool
}

func newPostNode(node *TreeNode, done bool) *PostNode {
	return &PostNode{node, done}
}

type postorderstack struct {
	ar []*PostNode
}

func newpostorderstack() *postorderstack {
	return &postorderstack{
		ar: make([]*PostNode, 0),
	}
}

func (s *postorderstack) push(node *PostNode) {
	s.ar = append(s.ar, node)
}

func (s *postorderstack) pop() *PostNode {
	if s.isempty() {
		return nil
	}
	node := s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return node
}

func (s *postorderstack) isempty() bool {
	return len(s.ar) == 0
}
