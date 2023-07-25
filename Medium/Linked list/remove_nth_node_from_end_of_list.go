package Linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{-1, head}

	prevOfRemoval, cur := dummyHead, dummyHead

	for cur.Next != nil {
		cur = cur.Next

		if n <= 0 {
			prevOfRemoval = prevOfRemoval.Next
		}

		n -= 1
	}

	prevOfRemoval.Next = prevOfRemoval.Next.Next
	return dummyHead.Next
}
