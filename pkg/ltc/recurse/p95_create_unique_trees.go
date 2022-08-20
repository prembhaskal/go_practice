package recurse

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	return genTreeRec(1, n)
}

// genTree(start, root)
// mid.Left = genTree(start, mid - 1)
// mid.Right = genTree(mid + 1, end)

func genTreeRec(start, end int) []*TreeNode {
	if start > end {
		return make([]*TreeNode, 0)
	}

	if start == end {
		node := &TreeNode{Val: start}
		return []*TreeNode{node}
	}

	alltrees := make([]*TreeNode, 0)

	for i := start; i <= end; i++ {

		lefttrees := genTreeRec(start, i-1)
		righttrees := genTreeRec(i+1, end)

		node := &TreeNode{Val: i}
		// alltrees = append(alltrees, cloneTree(node))

		for _, left := range lefttrees {
			for _, right := range righttrees {
				node.Left = left
				node.Right = right
				alltrees = append(alltrees, cloneTree(node))
			}
		}

		// no left, all right
		if len(lefttrees) == 0 {
			for _, right := range righttrees {
				node.Left = nil
				node.Right = right
				alltrees = append(alltrees, cloneTree(node))
			}
		}

		if len(righttrees) == 0 {
			for _, left := range lefttrees {
				node.Left = left
				node.Right = nil
				alltrees = append(alltrees, cloneTree(node))
			}
		}

	}

	return alltrees
}

func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newroot := &TreeNode{Val: root.Val}
	newleft := cloneTree(root.Left)
	newright := cloneTree(root.Right)

	newroot.Left = newleft
	newroot.Right = newright
	return newroot
}
