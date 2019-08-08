package SwapNode

import (
	"Algorithms"
)

func swapPairs(head *Algorithms.ListNode) *Algorithms.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, a, b *Algorithms.ListNode
	prev = new(Algorithms.ListNode)
	a, head = head, head.Next
	for a != nil && a.Next != nil {
		b = a.Next
		prev.Next, b.Next, a.Next, a, prev = b, a, b.Next, b.Next, a
	}
	return head
}
