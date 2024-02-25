package BST

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > val {
		root = searchBST(root.Left, val)
	} else if root.Val < val {
		root = searchBST(root.Right, val)
	}

	return root
}
