package DFS

import (
	"math"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isValidBST(root *TreeNode) bool {

	var valid func(node *TreeNode, left, right int) bool

	valid = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}

		if !(node.Val < right && node.Val > left) {
			return false
		}

		return valid(node.Left, left, node.Val) && valid(node.Right, node.Val, right)

	}

	return valid(root, math.MinInt64, math.MaxInt64)
}
