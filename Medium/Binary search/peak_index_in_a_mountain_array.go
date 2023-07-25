package Binary_search

func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-2

	for l <= r {
		mid := (l + r) / 2

		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}

		if arr[mid] < arr[mid-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
