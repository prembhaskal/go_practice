package bintree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	// root
	// left
	// right

	// -> root -> left (all below here... last_element) -> right (all below here...last_element)
	flattenlast(root)
	// return root
}

// return last node
func flattenlast(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	right := root.Right

	var leftlast *TreeNode
	if left != nil {
		leftlast = flattenlast(left)
		root.Right = left
	} else {
		leftlast = root
	}

	var rightlast *TreeNode
	if right != nil {
		rightlast = flattenlast(right)
		leftlast.Right = right
	} else {
		rightlast = leftlast
	}
	root.Left = nil
	return rightlast
}
