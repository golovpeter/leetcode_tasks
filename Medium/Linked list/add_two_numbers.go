package Linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res

	for l1 != nil || l2 != nil {
		if l1 != nil {
			temp.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}

		if temp.Val > 9 {
			temp.Val -= 10
			temp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			temp.Next = &ListNode{}
		}

		temp = temp.Next
	}

	return res
}
