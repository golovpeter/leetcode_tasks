package Linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge(left *ListNode, right *ListNode) *ListNode {
	res := &ListNode{}
	lastEl := res

	for left != nil || right != nil {
		if left == nil {
			lastEl.Next = right
			break
		}

		if right == nil {
			lastEl.Next = left
			break
		}

		if left.Val < right.Val {
			lastEl.Next = left
			lastEl = lastEl.Next
			left = left.Next
		} else {
			lastEl.Next = right
			lastEl = lastEl.Next
			right = right.Next
		}
	}

	return res.Next
}

func sortList(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	firstTail := slow
	slow = slow.Next
	firstTail.Next = nil

	first, second := sortList(head), sortList(slow)

	return merge(first, second)
}
