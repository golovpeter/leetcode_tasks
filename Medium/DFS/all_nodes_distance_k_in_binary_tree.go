package DFS

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := make(map[*TreeNode]*TreeNode)

	var dfs func(root, prev *TreeNode)

	// DFS
	dfs = func(root, prev *TreeNode) {
		if root == nil {
			return
		}

		parents[root] = prev
		dfs(root.Left, root)
		dfs(root.Right, root)
	}

	dfs(root, nil)

	var dist int
	var res []int

	queue := []*TreeNode{target}
	seen := make(map[*TreeNode]bool)
	seen[target] = true

	for len(queue) > 0 {
		if dist == k {
			for i := 0; i < len(queue); i++ {
				res = append(res, queue[i].Val)
			}

			return res
		}

		size := len(queue)

		for i := 0; i < size; i++ {
			if _, ok := seen[queue[i].Left]; !ok && queue[i].Left != nil {
				seen[queue[i].Left] = true
				queue = append(queue, queue[i].Left)
			}

			if _, ok := seen[queue[i].Right]; !ok && queue[i].Right != nil {
				seen[queue[i].Right] = true
				queue = append(queue, queue[i].Right)
			}

			if _, ok := seen[parents[queue[i]]]; !ok && parents[queue[i]] != nil {
				seen[parents[queue[i]]] = true
				queue = append(queue, parents[queue[i]])
			}
		}

		queue = queue[size:]
		dist++
	}

	return []int{}
}
