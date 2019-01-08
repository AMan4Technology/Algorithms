package lru

func Constructor(capacity int) LRUCache {
	var (
		head = &node{}
		tail = &node{}
	)
	head.next, tail.prev = tail, head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		data:     make(map[int]*node, capacity),
	}
}

type LRUCache struct {
	length, capacity int
	head, tail       *node
	data             map[int]*node
}

func (l *LRUCache) Get(key int) int {
	curr := l.get(key)
	if curr == nil {
		return -1
	}
	if curr != l.tail.prev {
		l.remove(curr)
		l.push(curr)
	}
	return curr.value
}

func (l *LRUCache) Put(key, value int) {
	curr := l.get(key)
	if curr == nil {
		if l.length == l.capacity {
			delete(l.data, l.head.next.key)
			l.length--
			l.remove(l.head.next)
		}
		node := &node{key: key, value: value}
		l.data[key] = node
		l.length++
		l.push(node)
	} else {
		curr.value = value
		if curr != l.tail.prev {
			l.remove(curr)
			l.push(curr)
		}
	}
}

func (l LRUCache) get(key int) *node {
	return l.data[key]
}

func (l LRUCache) remove(curr *node) {
	curr.prev.next, curr.next.prev = curr.next, curr.prev
}

func (l LRUCache) push(curr *node) {
	curr.prev, curr.next = l.tail.prev, l.tail
	l.tail.prev.next, l.tail.prev = curr, curr
}

type node struct {
	key, value int
	next, prev *node
}
