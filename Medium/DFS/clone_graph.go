package DFS

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	oldToNew := make(map[*Node]*Node)

	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		if _, ok := oldToNew[node]; ok {
			return oldToNew[node]
		}

		copyNode := &Node{Val: node.Val}
		oldToNew[node] = copyNode

		for _, el := range node.Neighbors {
			copyNode.Neighbors = append(copyNode.Neighbors, dfs(el))
		}

		return copyNode
	}

	return dfs(node)
}
