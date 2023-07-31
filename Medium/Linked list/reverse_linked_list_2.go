package Linked_list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	prev, cur := &ListNode{Val: -1, Next: head}, head

	for i := 0; i < left-1; i++ {
		prev, cur = prev.Next, cur.Next
	}

	var nxt *ListNode
	leftConnector, newTail := prev, cur

	for i := left; i <= right; i++ {
		nxt = cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}

	newTail.Next = nxt

	if leftConnector.Next != head {
		leftConnector.Next = prev
	} else {
		head = prev
	}

	return head
}
