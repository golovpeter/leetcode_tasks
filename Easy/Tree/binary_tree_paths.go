package Tree

import (
	"strconv"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func binaryTreePaths(head *TreeNode) []string {
	var res []string
	var dfs func(*TreeNode, string)

	dfs = func(node *TreeNode, curSeq string) {
		if node == nil {
			return
		}

		if len(curSeq) == 0 {
			curSeq += strconv.Itoa(node.Val)
		} else {
			curSeq += "->" + strconv.Itoa(node.Val)
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, curSeq)
			return
		}

		dfs(node.Left, curSeq)
		dfs(node.Right, curSeq)
	}

	dfs(head, "")
	return res
}
