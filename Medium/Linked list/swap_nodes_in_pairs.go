package Linked_list

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev, cur := dummyHead, head

	for cur != nil && cur.Next != nil {
		nextPair := cur.Next.Next
		second := cur.Next

		second.Next = cur
		cur.Next = nextPair
		prev.Next = second

		prev = cur
		cur = nextPair
	}

	return dummyHead.Next
}
