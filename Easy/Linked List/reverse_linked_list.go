package Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, prev, nxt *ListNode
	cur = head

	for cur != nil {
		nxt = cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}

	return prev
}
