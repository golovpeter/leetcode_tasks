package Two_pointers

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]

			if target == tmp {
				return target
			}

			if abs(target-res) > abs(target-tmp) {
				res = tmp
			}

			if tmp > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
