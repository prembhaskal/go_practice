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
	// valid, _, _ := checkMinMax(root)
	// return valid
	isvalid, _ := inorderBSTCheck(root, nil)
	return isvalid
}

func inorderBSTCheck(root, prev *TreeNode) (bool, *TreeNode) {
	if root == nil {
		return true, prev
	}
	// if prev != nil {
	//     fmt.Printf("root: %d, prev: %d\n", root.Val, prev.Val)
	// } else {
	//     fmt.Printf("root: %d, prev: nil\n", root.Val)
	// }
	var isvalid bool
	// if root.Left != nil {
	isvalid, prev = inorderBSTCheck(root.Left, prev)
	if !isvalid {
		return false, prev
	}
	// }
	// if prev != nil {
	//     fmt.Printf("curr: %d, prev: %d\n", root.Val, prev.Val)
	// }
	if prev != nil && prev.Val >= root.Val {
		return false, prev
	}

	return inorderBSTCheck(root.Right, root)
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
