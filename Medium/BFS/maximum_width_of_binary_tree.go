package BFS

type node struct {
	val *TreeNode
	num int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	queue := []*node{{root, 0}}

	for {
		var tmp []*node

		for i := 0; i < len(queue); i++ {
			if queue[i].val.Left != nil {
				tmp = append(tmp, &node{queue[i].val.Left, queue[i].num * 2})
			}

			if queue[i].val.Right != nil {
				tmp = append(tmp, &node{queue[i].val.Right, queue[i].num*2 + 1})
			}
		}

		if len(tmp) > 0 {
			w := tmp[len(tmp)-1].num - tmp[0].num

			if w > res {
				res = w
			}
		} else {
			break
		}

		queue = tmp
	}

	return res
}
