package Tree

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	if s.Val == t.Val && isIdentical(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isIdentical(s, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return s.Val == t.Val && isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
}
