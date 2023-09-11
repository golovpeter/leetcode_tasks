package Array

func merge(nums1 []int, m int, nums2 []int, n int) {
	var ptr1, ptr2, last int = m - 1, n - 1, m + n - 1

	for ptr1 >= 0 && ptr2 >= 0 {
		if nums1[ptr1] > nums2[ptr2] {
			nums1[last] = nums1[ptr1]
			ptr1--
		} else {
			nums1[last] = nums2[ptr2]
			ptr2--
		}

		last--
	}

	for ptr2 >= 0 {
		nums1[last] = nums2[ptr2]
		ptr2--
		last--
	}
}
