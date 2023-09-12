package DFS

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	} else {
		return r
	}
}
