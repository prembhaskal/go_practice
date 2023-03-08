package dp

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

var (
	dp1 map[*TreeNode]int
	dp2 map[*TreeNode]int
)

func rob3(root *TreeNode) int {
	dp1 = make(map[*TreeNode]int)
	dp2 = make(map[*TreeNode]int)
	return max337(rob31(root, true), rob31(root, false))
}

// choose root
// cannot choose child, choose next level

// don't choose root,
// choose(children)

func rob31(root *TreeNode, skip bool) int {
	if root == nil {
		return 0
	}

	if skip {
		if val, ok := dp2[root]; ok {
			return val
		}
	} else {
		if val, ok := dp1[root]; ok {
			return val
		}
	}
	left := root.Left
	right := root.Right

	if skip {
		dp2[root] = rob31(left, false) + rob31(right, false)
		return dp2[root]
	}

	// choose
	v1 := rob31(left, true) + rob31(right, true) + root.Val
	// don't choose
	v2 := rob31(left, false) + rob31(right, false)
	dp1[root] = max337(v1, v2)
	return dp1[root]
}

func max337(a, b int) int {
	if a > b {
		return a
	}
	return b
}
