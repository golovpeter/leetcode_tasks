package Heap

import (
	"container/heap"
	"strings"
)

func frequencySort(s string) string {
	freq := make(map[rune]int)
	h := &LetterHeap{}
	heap.Init(h)

	for _, val := range s {
		freq[val]++
	}

	for key, val := range freq {
		heap.Push(h, LetterPair{val, key})
	}

	var res strings.Builder

	for h.Len() > 0 {
		el := heap.Pop(h).(LetterPair)

		for i := 0; i < el.Freq; i++ {
			res.WriteRune(el.Value)
		}
	}

	return res.String()
}

type LetterPair struct {
	Freq  int
	Value rune
}

type LetterHeap []LetterPair

func (h LetterHeap) Len() int {
	return len(h)
}

func (h LetterHeap) Less(i, j int) bool {
	return h[i].Freq > h[j].Freq
}

func (h LetterHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *LetterHeap) Push(x any) {
	*h = append(*h, x.(LetterPair))
}

func (h *LetterHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
