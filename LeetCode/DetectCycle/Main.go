package DetectCycle

import (
	"Algorithms"
)

func hasCycle(head *Algorithms.ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
