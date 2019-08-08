package KClosestPointsToOrigin

func kClosest(points [][]int, K int) (results [][]int) {
	queue := NewMax(K)
	for i, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		if !queue.Full() {
			queue.Enqueue(i, distance)
			continue
		}
		top := queue.Top()
		if distance < top.Value {
			queue.Update(i, distance)
		}
	}
	results = make([][]int, K)
	queue.Range(func(i int, node Node) bool {
		results[K-i] = points[node.ID]
		return true
	})
	return
}

func NewMax(capacity int) Max {
	return Max{
		capacity: capacity,
		data:     make([]Node, 0, capacity),
	}
}

type Max struct {
	capacity, length int
	data             []Node
}

func (h Max) Full() bool {
	return h.length == h.capacity
}

func (h Max) Empty() bool {
	return h.length == 0
}

func (h *Max) Enqueue(id, value int) bool {
	if h.Full() {
		return false
	}
	h.length++
	h.data = append(h.data, Node{id, value})
	h.up()
	return true
}

func (h *Max) Dequeue() (top Node) {
	if h.length == 0 {
		return Node{}
	}
	top, h.data[0] = h.data[0], h.data[h.length-1]
	h.length--
	h.data = h.data[:h.length]
	h.down()
	return
}

func (h Max) Top() Node {
	if h.length == 0 {
		return Node{}
	}
	return h.data[0]
}

func (h Max) Update(id, value int) {
	h.data[0].ID, h.data[0].Value = id, value
	h.down()
}

func (h Max) Range(callback func(i int, node Node) bool) {
	for i, node := range h.data {
		if !callback(i, node) {
			break
		}
	}
}

func (h Max) up() {
	for curr := h.length - 1; h.data[(curr-1)/2].Value < h.data[curr].Value; curr = (curr - 1) / 2 {
		h.data[(curr-1)/2], h.data[curr] = h.data[curr], h.data[(curr-1)/2]
	}
}

func (h Max) down() {
	for curr, max := 0, 0; ; curr = max {
		if 2*curr+1 < h.length && h.data[2*curr+1].Value > h.data[max].Value {
			max = 2*curr + 1
		}
		if 2*curr+2 < h.length && h.data[2*curr+2].Value > h.data[max].Value {
			max = 2*curr + 2
		}
		if max == curr {
			break
		}
		h.data[max], h.data[curr] = h.data[curr], h.data[max]
	}
}

type Node struct {
	ID, Value int
}
