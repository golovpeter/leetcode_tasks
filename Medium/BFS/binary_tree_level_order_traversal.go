package BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{{root.Val}}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n := len(queue)
		var level []int

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				level = append(level, node.Left.Val)
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				level = append(level, node.Right.Val)
				queue = append(queue, node.Right)
			}
		}

		if len(level) != 0 {
			res = append(res, level)
		}
	}

	return res
}
