package heap

func NewMin(capacity int) Min {
	return Min{
		capacity: capacity,
		data:     make([]Node, 0, capacity),
	}
}

type Min struct {
	capacity, length int
	data             []Node
}

func (h Min) Full() bool {
	return h.length == h.capacity
}

func (h Min) Empty() bool {
	return h.length == 0
}

func (h *Min) Push(id, value int) bool {
	if h.Full() {
		return false
	}
	h.length++
	h.data = append(h.data, Node{id, value})
	h.up()
	return true
}

func (h *Min) Pop() (top Node) {
	if h.length == 0 {
		return Node{}
	}
	top, h.data[0] = h.data[0], h.data[h.length-1]
	h.length--
	h.data = h.data[:h.length]
	h.down()
	return
}

func (h Min) Top() Node {
	if h.length == 0 {
		return Node{}
	}
	return h.data[0]
}

func (h Min) Update(value int) {
	h.data[0].Value = value
	h.down()
}

func (h Min) up() {
	for curr := h.length - 1; h.data[(curr-1)/2].Value > h.data[curr].Value; curr = (curr - 1) / 2 {
		h.data[(curr-1)/2], h.data[curr] = h.data[curr], h.data[(curr-1)/2]
	}
}

func (h Min) down() {
	for curr, min := 0, 0; ; curr = min {
		if 2*curr+1 < h.length && h.data[2*curr+1].Value < h.data[min].Value {
			min = 2*curr + 1
		}
		if 2*curr+2 < h.length && h.data[2*curr+2].Value < h.data[min].Value {
			min = 2*curr + 2
		}
		if min == curr {
			break
		}
		h.data[min], h.data[curr] = h.data[curr], h.data[min]
	}
}

type Node struct {
	ID, Value int
}
