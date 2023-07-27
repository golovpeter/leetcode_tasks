package Tree

func minDepth(root *TreeNode) int {
	queue := []*TreeNode{root}
	minDepth := 1

	if root == nil {
		return 0
	}

	for len(queue) != 0 {
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return minDepth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		minDepth++
	}

	return minDepth
}
