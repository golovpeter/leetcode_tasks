package Linked_list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	count, lastEl := 1, head
	for lastEl.Next != nil {
		lastEl = lastEl.Next
		count++
	}

	prev, cur := head, head.Next

	for i := 2; i <= count; i++ {
		if i%2 == 0 {
			nxt := cur.Next
			prev.Next = cur.Next
			lastEl.Next = cur
			cur.Next = nil
			cur = nxt
			lastEl = lastEl.Next
		} else {
			cur = cur.Next
			prev = prev.Next
		}
	}

	return head
}
