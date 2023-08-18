package DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var n, result int

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		n++

		if n == k {
			result = node.Val
			return
		}

		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return result
}
