package Sliding_window

func numSubarrayProductLessThanK(nums []int, k int) int {
	var l, result int
	curr := 1

	for r := 0; r < len(nums); r++ {
		curr *= nums[r]

		for l <= r && curr >= k {
			curr /= nums[l]
			l++
		}

		result += r - l + 1
	}

	return result
}
