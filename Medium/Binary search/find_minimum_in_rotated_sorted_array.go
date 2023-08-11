package Binary_search

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
