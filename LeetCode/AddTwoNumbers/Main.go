package AddTwoNumbers

import (
	"Algorithms"
)

func addTwoNumbers(l1 *Algorithms.ListNode, l2 *Algorithms.ListNode) (head *Algorithms.ListNode) {
	if l1 == nil && l2 == nil {
		return nil
	}
	head = new(Algorithms.ListNode)
	var (
		prev = head
		curr = new(Algorithms.ListNode)
	)
	for l1 != nil || l2 != nil || curr.Value != 0 {
		if l1 != nil {
			curr.Value += l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			curr.Value += l2.Value
			l2 = l2.Next
		}
		prev.Next = curr
		curr.Value, prev, curr = curr.Value%10, curr, &Algorithms.ListNode{Value: curr.Value / 10}
	}
	return head.Next
}
