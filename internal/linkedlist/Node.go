package linkedlist

import (
	"Algorithms"
)

func mergeKLists(lists []*ListNode) (head *ListNode) {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	heap := Algorithms.NewMin(length)
	currents := make([]*ListNode, 0, length)
	for _, curr := range lists {
		if curr != nil {
			currents = append(currents, curr)
			heap.Push(len(currents)-1, curr.Value)
		}
	}
	head = new(ListNode)
	for prev := head; !heap.Empty(); prev = prev.Next {
		curr := heap.Top()
		prev.Next = &ListNode{Value: curr.Value}
		currents[curr.ID] = currents[curr.ID].Next
		if currents[curr.ID] != nil {
			heap.Update(currents[curr.ID].Value)
		} else {
			heap.Pop()
		}
	}
	return head.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head = new(ListNode)
	for i, j, prev := l1, l2, head; i != nil || j != nil; prev = prev.Next {
		if i == nil {
			prev.Next = &ListNode{Value: j.Value}
			j = j.Next
			continue
		}
		if j == nil {
			prev.Next = &ListNode{Value: i.Value}
			i = i.Next
			continue
		}
		if i.Value <= j.Value {
			prev.Next = &ListNode{Value: i.Value}
			i = i.Next
		} else {
			prev.Next = &ListNode{Value: j.Value}
			j = j.Next
		}
	}
	return head.Next
}

type ListNode struct {
	Value int
	Next  *ListNode
}
