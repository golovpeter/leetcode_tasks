package BFS

type nod struct {
	Val   int
	Left  *nod
	Right *nod
	Next  *nod
}

func connect(root *nod) *nod {
	if root == nil {
		return nil
	}

	queue := []*nod{root}

	for len(queue) != 0 {
		n := len(queue)

		var lastNode *nod

		for i := 0; i < n; i++ {
			lastNode = queue[0]
			queue = queue[1:]

			if len(queue) != 0 {
				lastNode.Next = queue[0]
			}

			if lastNode.Left != nil {
				queue = append(queue, lastNode.Left)
			}

			if lastNode.Right != nil {
				queue = append(queue, lastNode.Right)
			}
		}

		lastNode.Next = nil
	}

	return root
}
