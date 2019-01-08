package ReverseLinkedList

import . "Algorithms/LeetCode/internal/linkedlist"

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, cur *ListNode
	cur = head
	for cur != nil {
		cur.Next, cur, prev = prev, cur.Next, cur
	}
	return prev
}
