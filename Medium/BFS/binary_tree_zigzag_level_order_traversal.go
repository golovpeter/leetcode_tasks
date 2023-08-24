package BFS

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}
	flag := true

	for len(queue) != 0 {
		n := len(queue)
		var level []int

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if flag {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
		}

		flag = !flag
		res = append(res, level)
	}

	return res
}
