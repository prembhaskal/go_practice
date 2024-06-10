package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	kthElem = -1

	inOrder(root, k)

	return kthElem
}

var kthElem int

// inOrder visit tree, reducing k for each visit
// returns count of visited nodes
func inOrder(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	left := inOrder(root.Left, k)

	if k-left == 1 {
		kthElem = root.Val
	}

	right := inOrder(root.Right, k-left-1)

	return left + right + 1
}
