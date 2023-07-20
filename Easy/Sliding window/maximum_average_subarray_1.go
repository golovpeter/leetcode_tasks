package Sliding_window

func findMaxAverage(nums []int, k int) float64 {
	var prevSum int
	left, right := 0, k-1

	for i := left; i <= right; i++ {
		prevSum += nums[i]
	}

	maxAverage := float64(prevSum) / float64(k)
	right++
	left++

	for right <= len(nums)-1 {
		prevSum -= nums[left-1]
		prevSum += nums[right]

		curAverage := float64(prevSum) / float64(k)
		if curAverage > maxAverage {
			maxAverage = curAverage
		}

		left++
		right++
	}

	return maxAverage
}
