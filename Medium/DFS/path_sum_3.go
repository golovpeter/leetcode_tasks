package DFS

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func pathSum3(root *TreeNode, targetSum int) int {
	return dfs(root, targetSum, []int{})
}

func dfs(root *TreeNode, targetSum int, currentPath []int) int {
	if root == nil {
		return 0
	}

	newCurrPath := make([]int, len(currentPath))
	copy(newCurrPath, currentPath)

	newCurrPath = append(newCurrPath, root.Val)
	pathCount, pathSum := 0, 0

	for i := len(newCurrPath); i >= 0; i-- {
		pathSum += newCurrPath[i]

		if pathSum == targetSum {
			pathCount++
		}
	}

	pathCount += dfs(root.Left, targetSum, newCurrPath)
	pathCount += dfs(root.Right, targetSum, newCurrPath)

	return pathCount
}
