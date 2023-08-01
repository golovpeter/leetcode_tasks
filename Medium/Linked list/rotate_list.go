package Linked_list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	count, tmp := 0, head
	dummy := &ListNode{Next: head}

	for tmp.Next != nil {
		tmp = tmp.Next
		count++
	}

	count++
	j := k % count
	sp := head

	for i := 0; i < j; i++ {
		sp = sp.Next
	}

	tmp.Next = head
	dummy.Next = sp.Next
	sp.Next = nil

	return dummy.Next
}
