package Heap

import (
	"container/heap"
)

type Set struct {
	Sum  int
	i, j int
}

type SetHeap []Set

func (h SetHeap) Len() int {
	return len(h)
}

func (h SetHeap) Less(i, j int) bool {
	return h[i].Sum < h[j].Sum
}

func (h SetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SetHeap) Push(x any) {
	*h = append(*h, x.(Set))
}

func (h *SetHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result [][]int
	minHeap := &SetHeap{}
	heap.Init(minHeap)

	for i, val := range nums1 {
		heap.Push(minHeap, Set{val + nums2[0], i, 0})
	}

	for minHeap.Len() > 0 && k > 0 {
		minSum := heap.Pop(minHeap).(Set)

		currI := minSum.i
		currJ := minSum.j

		result = append(result, []int{nums1[currI], nums2[currJ]})

		if currJ+1 < len(nums2) {
			heap.Push(minHeap, Set{nums1[currI] + nums2[currJ+1], currI, currJ + 1})
		}

		k--
	}

	return result
}
