package SwapNode

import . "Algorithms/LeetCode/internal/linkedlist"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, a, b *ListNode
	prev = new(ListNode)
	a, head = head, head.Next
	for a != nil && a.Next != nil {
		b = a.Next
		prev.Next, b.Next, a.Next, a, prev = b, a, b.Next, b.Next, a
	}
	return head
}
