package Linked_list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func partition(head *ListNode, x int) *ListNode {
	before, after := &ListNode{Val: -1}, &ListNode{Val: -1}
	bTail, aTail := before, after

	for head != nil {
		if head.Val < x {
			bTail.Next = head
			bTail = bTail.Next

		} else {
			aTail.Next = head
			aTail = aTail.Next
		}

		head = head.Next
	}

	aTail.Next = nil
	bTail.Next = after.Next

	return before.Next
}
