package Arrays

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	var res int

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i] = intervals[i-1]
			res++
		}
	}

	return res
}
