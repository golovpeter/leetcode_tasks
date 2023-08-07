package Arrays

import (
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for _, val := range intervals[1:] {
		lastEnd := result[len(result)-1][1]

		if val[0] <= lastEnd {
			result[len(result)-1][1] = max(val[1], lastEnd)
		} else {
			result = append(result, val)
		}
	}

	return result
}
