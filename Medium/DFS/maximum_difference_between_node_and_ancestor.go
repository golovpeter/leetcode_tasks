package DFS

func maxAncestorDiff(root *TreeNode) int {
	var res int
	curMax, curMin := root.Val, root.Val

	var dfs func(node *TreeNode, curMax, curMin int)
	dfs = func(node *TreeNode, curMax, curMin int) {
		if node == nil {
			return
		}

		if node.Val > curMax {
			curMax = node.Val
		}

		if node.Val < curMin {
			curMin = node.Val
		}

		if res < abc(curMin-curMax) {
			res = abc(curMin - curMax)
		}

		dfs(node.Left, curMax, curMin)
		dfs(node.Right, curMax, curMin)
	}

	dfs(root, curMax, curMin)

	return res
}

func abc(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}
