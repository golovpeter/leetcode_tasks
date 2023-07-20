package Backtracking

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	res := make([][]int, 0)
	visited := make(map[int]bool)
	list := make([]int, 0)
	backTrack(nums, visited, list, &res)
	return res
}

func backTrack(nums []int, visited map[int]bool, list []int, res *[][]int) {
	//back track goal, length of list equals to length of nums
	if len(list) == len(nums) {
		temp := make([]int, len(list))
		copy(temp, list) //copy list into temp, so it won't impact subsequent process on list
		*res = append(*res, temp)
		return
	}

	//options are each elements in nums slice
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue //constrain, ignore visited element
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue //constrain, if there is a duplication, the first element need to be visited first
		}
		list = append(list, nums[i])
		visited[i] = true
		backTrack(nums, visited, list, res)
		list = list[:len(list)-1]
		visited[i] = false
	}
}
