package Arrays

func rotate(nums []int, k int) {
	res := make([]int, len(nums))

	copy(res, nums)

	for i, num := range res {
		g := i + k

		if g >= len(res)-1 {
			g = (i + k) % len(res)
		}

		nums[g] = num
	}
}
