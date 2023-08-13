package Binary_search

func search2(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if target == nums[mid] {
			return true
		}

		if nums[mid] > nums[l] {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[l] > nums[mid] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			l++
		}
	}

	return false
}
