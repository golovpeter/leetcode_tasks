package DFS

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var dfs func(root *TreeNode, targetSum int, currComb []int)

	dfs = func(root *TreeNode, targetSum int, currComb []int) {
		if root == nil {
			return
		}

		targetSum -= root.Val
		currComb = append(currComb, root.Val)

		if root.Right == nil && root.Left == nil {

			if targetSum == 0 {
				copyCurComb := make([]int, len(currComb))
				copy(copyCurComb, currComb)
				res = append(res, copyCurComb)
				return
			}
		}

		dfs(root.Left, targetSum, currComb)
		dfs(root.Right, targetSum, currComb)
	}

	dfs(root, targetSum, []int{})
	return res
}
