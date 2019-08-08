package ReverseLinkedList

import (
	"Algorithms"
)

func reverseList(head *Algorithms.ListNode) *Algorithms.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, cur *Algorithms.ListNode
	cur = head
	for cur != nil {
		cur.Next, cur, prev = prev, cur.Next, cur
	}
	return prev
}
