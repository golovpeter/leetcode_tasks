package Arrays

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//
//	return b
//}

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else if newInterval[0] > intervals[i][1] {
			result = append(result, intervals[i])
		} else {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
		}
	}

	result = append(result, newInterval)
	return result
}
