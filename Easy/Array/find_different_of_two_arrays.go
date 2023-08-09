package Array

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)

	for _, val := range nums1 {
		map1[val] = true
	}

	map2 := make(map[int]bool)

	for _, val := range nums2 {
		map2[val] = true
	}

	var ans1, ans2 []int

	for _, val := range nums1 {
		if !map2[val] {
			ans1 = append(ans1, val)
			map2[val] = true
		}
	}

	for _, val := range nums2 {
		if !map1[val] {
			ans2 = append(ans2, val)
			delete(map1, val)
			map1[val] = true
		}
	}

	return [][]int{ans1, ans2}
}
