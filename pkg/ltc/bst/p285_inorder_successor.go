package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Right != nil {
		return treeminimum(p.Right)
	}

	// search from top to down
	curr := root
	// left and right parent
	var lp *TreeNode

	for curr != nil {
		if curr == p {
			return lp
		}

		if curr.Val > p.Val {
			lp = curr
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}

func treeminimum(r *TreeNode) *TreeNode {
	for r.Left != nil {
		r = r.Left
	}
	return r
}
