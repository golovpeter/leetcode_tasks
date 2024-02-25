package BFS

import "math"

type Edge struct {
	Node int
	Cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]Edge, n)

	for _, el := range flights {
		graph[el[0]] = append(graph[el[0]], Edge{el[1], el[2]})
	}

	queue := []Edge{{src, 0}}

	minCost := math.MaxInt32

	costs := make([]int, n)
	for i := range costs {
		costs[i] = math.MaxInt32
	}

	costs[src] = 0
	for k >= 0 && len(queue) > 0 {
		nextQueue := []Edge{}
		nextCosts := append([]int{}, costs...)

		for _, trip := range queue {
			if costs[trip.Node] != trip.Cost {
				continue
			}

			for _, edge := range graph[trip.Node] {
				nextCost := trip.Cost + edge.Cost

				if edge.Node == dst {
					minCost = min(minCost, nextCost)
					continue
				}

				nextCosts[edge.Node] = min(nextCosts[edge.Node], nextCost)
				nextQueue = append(nextQueue, Edge{edge.Node, nextCost})
			}
		}

		queue = nextQueue
		costs = nextCosts
		k--
	}

	if minCost == math.MaxInt32 {
		return -1
	}

	return minCost
}
