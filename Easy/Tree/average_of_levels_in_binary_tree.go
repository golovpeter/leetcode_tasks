package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n := len(queue)
		levelSum := 0.0

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, levelSum/float64(n))
	}

	return res
}
