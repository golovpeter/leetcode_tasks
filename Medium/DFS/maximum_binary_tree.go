package DFS

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func getMaxInArray(arr []int) (int, int) {
	idx, value := 0, 0

	for i, v := range arr {
		if v > value {
			value = v
			idx = i
		}
	}

	return idx, value
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	i, v := getMaxInArray(nums)
	return &TreeNode{
		Val:   v,
		Left:  constructMaximumBinaryTree(nums[:i]),
		Right: constructMaximumBinaryTree(nums[i+1:]),
	}
}
