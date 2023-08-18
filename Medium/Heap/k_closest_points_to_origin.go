package Heap

import (
	"container/heap"
)

func kClosest(points [][]int, k int) [][]int {
	var res [][]int

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for i, val := range points {
		heap.Push(minHeap, Pair{
			Distance: (val[0] * val[0]) + (val[1] * val[1]),
			Idx:      i,
		})
	}

	for i := 0; i < k; i++ {
		lowEL := heap.Pop(minHeap).(Pair)
		res = append(res, points[lowEL.Idx])
	}

	return res
}

type Pair struct {
	Distance int
	Idx      int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
