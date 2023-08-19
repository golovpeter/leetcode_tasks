package Heap

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	h := &MinHeapPair{}
	heap.Init(h)

	for _, val := range nums {
		freq[val]++
	}

	for key, val := range freq {
		heap.Push(h, FreqPair{val, key})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)

	for i := 1; i <= k; i++ {
		res[k-i] = heap.Pop(h).(FreqPair).Value
	}

	return res
}

type FreqPair struct {
	Freq  int
	Value int
}

type MinHeapPair []FreqPair

func (h MinHeapPair) Len() int {
	return len(h)
}

func (h MinHeapPair) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h MinHeapPair) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapPair) Push(x any) {
	*h = append(*h, x.(FreqPair))
}

func (h *MinHeapPair) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
