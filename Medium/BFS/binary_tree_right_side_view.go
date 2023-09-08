package BFS

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func rightSideView(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	var prev *TreeNode

	for len(queue) != 0 {
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			prev = node
		}

		res = append(res, prev.Val)
	}

	return res
}
