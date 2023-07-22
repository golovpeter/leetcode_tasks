package Linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	visited[head] = true

	if head == nil || head.Next == nil {
		return nil
	}

	fast, slow := head.Next, head

	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		visited[fast] = true
		fast = fast.Next

		if visited[fast] {
			return fast
		}

	}
}
