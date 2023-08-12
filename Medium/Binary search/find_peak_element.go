package Binary_search

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	if len(nums) >= 2 {
		if nums[r] > nums[r-1] {
			return r
		} else if nums[l] > nums[l+1] {
			return l
		}
	}

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid] < nums[mid-1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return 0
}
