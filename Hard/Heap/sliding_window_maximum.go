package Heap

import "container/heap"

type Pair struct {
	index int
	value int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].value > h[j].value || (h[i].value == h[j].value && h[i].index < h[j].index)
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)

	h := &PairHeap{}
	heap.Init(h)

	l, r, idx := 0, 0, 0

	for r < len(nums) {
		pair := Pair{r, nums[r]}
		heap.Push(h, pair)

		if r-l+1 < k {
			r++
		} else if r-l+1 == k {
			for (*h)[0].index < l {
				heap.Pop(h)
			}

			first := (*h)[0]
			result[idx] = first.value

			if first.index < l+1 {
				heap.Pop(h)
			}

			l++
			r++
			idx++
		}
	}

	return result
}
