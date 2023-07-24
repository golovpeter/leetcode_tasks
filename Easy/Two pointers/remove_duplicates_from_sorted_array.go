package Two_pointers

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return 1
	}

	l := 0

	for i := range nums {
		if nums[l] != nums[i] {
			l++
			nums[l] = nums[i]
		}
	}

	return l + 1
}
