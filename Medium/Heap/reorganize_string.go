package Heap

import (
	"container/heap"
	"strings"
)

func reorganizeString(s string) string {
	var res strings.Builder
	freq := make(map[rune]int)

	for _, val := range s {
		freq[val]++
	}

	h := &MaxHeap{}
	heap.Init(h)

	for key, val := range freq {
		heap.Push(h, &LetPair{key, val})
	}

	var prev *LetPair
	for h.Len() != 0 || prev != nil {
		if prev != nil && h.Len() == 0 {
			return ""
		}

		topEl := heap.Pop(h).(*LetPair)
		res.WriteRune(topEl.Value)
		topEl.Freq--

		if prev != nil {
			heap.Push(h, prev)
			prev = nil
		}

		if topEl.Freq != 0 {
			prev = topEl
		}
	}

	return res.String()
}

type LetPair struct {
	Value rune
	Freq  int
}

type MaxHeap []*LetPair

func (p MaxHeap) Len() int           { return len(p) }
func (p MaxHeap) Less(i, j int) bool { return p[i].Freq > p[j].Freq }
func (p MaxHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *MaxHeap) Push(x any)        { *p = append(*p, x.(*LetPair)) }

func (p *MaxHeap) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
