package Arrays

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int

	if len(firstList) == 0 || len(secondList) == 0 {
		return result
	}

	l1, l2 := 0, 0

	for l1 < len(firstList) && l2 < len(secondList) {
		firstListItem := firstList[l1]
		secondListItem := secondList[l2]

		if firstListItem[1] < secondListItem[0] {
			l1++
			continue
		}

		if secondListItem[1] < firstListItem[0] {
			l2++
			continue
		}

		low := max(firstListItem[0], secondListItem[0])
		high := min(firstListItem[1], secondListItem[1])

		result = append(result, []int{low, high})

		if firstListItem[1] > secondListItem[1] {
			l2++
		} else {
			l1++
		}
	}

	return result
}
