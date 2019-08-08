package OddEvenJump

func oddEvenJumps(A []int) (count int) {
	length := len(A)
	if length < 2 {
		return length
	}
	var (
		odds  = make([]bool, length)
		evens = make([]bool, length)
		tree  = newNode(length-1, A[length-1])
	)
	odds[length-1], evens[length-1] = true, true
	count = 1
	for i := length - 2; i >= 0; i-- {
		prev, next := tree.Save(i, A[i])
		if prev != nil {
			var index int
			if A[i] == prev.Val {
				index = prev.indexes[len(prev.indexes)-2]
			} else {
				index = prev.indexes[len(prev.indexes)-1]
			}
			if odds[index] {
				evens[i] = true
			}
		}
		if next != nil {
			var index int
			if A[i] == next.Val {
				index = next.indexes[len(next.indexes)-2]
			} else {
				index = next.indexes[len(next.indexes)-1]
			}
			if evens[index] {
				odds[i] = true
				count++
			}
		}
	}
	return
}

func newNode(index, value int) *node {
	return &node{
		Val:     value,
		indexes: []int{index}}
}

type node struct {
	Val         int
	Left, Right *node
	indexes     []int
}

func (n *node) Save(index, value int) (prev, next *node) {
	for curr := n; ; {
		if value < curr.Val {
			next = curr
			if curr.Left == nil {
				curr.Left = newNode(index, value)
				return
			}
			curr = curr.Left
		} else if curr.Val < value {
			prev = curr
			if curr.Right == nil {
				curr.Right = newNode(index, value)
				return
			}
			curr = curr.Right
		} else {
			curr.indexes = append(curr.indexes, index)
			return curr, curr
		}
	}
}
