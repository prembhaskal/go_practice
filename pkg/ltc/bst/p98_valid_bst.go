package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	valid, _, _ := checkMinMax(root)
	return valid
}

func checkMinMax(root *TreeNode) (bool, int, int) {
	var lmin, lmax, rmin, rmax int
	var lvalid, rvalid bool
	if root.Left != nil {
		lvalid, lmin, lmax = checkMinMax(root.Left)
		if !lvalid {
			return false, -1, -1 // for false, values don't matter
		}
		if root.Val <= lmax {
			return false, -1, -1
		}
	}
	if root.Right != nil {
		rvalid, rmin, rmax = checkMinMax(root.Right)
		if !rvalid || root.Val >= rmin {
			return false, -1, -1
		}
	}

	if root.Left == nil {
		if root.Right == nil {
			return true, root.Val, root.Val
		}
		return true, min(root.Val, rmin), max(root.Val, rmax)
	}
	if root.Right == nil {
		// fmt.Printf("lmin:%d, lmax:%d\n", lmin, lmax)
		return true, min(root.Val, lmin), max(root.Val, lmax)
	}

	return true, min(root.Val, min(lmin, rmin)), max(root.Val, max(rmin, rmax))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
