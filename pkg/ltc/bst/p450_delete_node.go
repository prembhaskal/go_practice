package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	node, parent := getNodeAndParent(root, key)
	if node == nil {
		return root
	}
	if parent == nil {
		return removeRoot(root)
	}
	// fmt.Printf("node: %d, parent: %d\n", node.Val, parent.Val)
	if node.Right == nil {
		transplant(node, parent, node.Left)
	} else {
		next, nextpr := getTreeMinimum(node.Right, node)
		// fmt.Printf("next: %d, nextparent: %d\n", next.Val, nextpr.Val)
		// replace node with next
		if nextpr != node { // special case, to avoid cycle
			// earlier nextpr.Left was pointing to next, make it point to next.right
			// nextpr.Left = next.Right
			transplant(next, nextpr, next.Right)
			next.Right = node.Right
		}

		transplant(node, parent, next)
		next.Left = node.Left
	}

	node.Left = nil
	node.Right = nil

	return root
}

// from CLRS, replace subtree rooted at old with nnew
func transplant(old, oldp, nnew *TreeNode) {
	if oldp.Left == old {
		oldp.Left = nnew
	} else {
		oldp.Right = nnew
	}
}

func removeRoot(root *TreeNode) *TreeNode {
	if root.Right == nil {
		// if root.Left == nil {
		//     return nil
		// }
		return root.Left
	}
	next, nextpr := getTreeMinimum(root.Right, root)
	if root.Right != next { // avoid cycle if root is parent of next
		next.Right = root.Right // make old roots right
	}
	next.Left = root.Left
	// remove next from tree
	transplant(next, nextpr, nil)
	root.Right = nil // remove reference
	return next
}

// get node and parent
func getNodeAndParent(root *TreeNode, key int) (*TreeNode, *TreeNode) {
	par := (*TreeNode)(nil)
	next := root
	for next != nil {
		if key < next.Val {
			par = next
			next = next.Left
		} else if key > next.Val {
			par = next
			next = next.Right
		} else {
			return next, par
		}
	}
	return nil, nil
}

// get min node and its parent (possibly nil)
func getTreeMinimum(node, prev *TreeNode) (*TreeNode, *TreeNode) {
	// prev := nil
	next := node
	for next.Left != nil {
		prev = next
		next = next.Left
	}

	return next, prev
}
