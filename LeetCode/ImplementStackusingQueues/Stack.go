package ImplementStackusingQueues

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0, 5),
		temp:  make([]int, 0, 5),
	}
}

type MyStack struct {
	queue, temp []int
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	for i := 0; i < len(this.queue)-1; i++ {
		this.temp = append(this.temp, this.queue[i])
	}
	this.queue = this.queue[len(this.queue)-1:]
	defer func() {
		this.queue = this.queue[:0]
		this.queue, this.temp = this.temp, this.queue
	}()
	return this.queue[0]
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	for i := 0; i < len(this.queue)-1; i++ {
		this.temp = append(this.temp, this.queue[i])
	}
	this.queue = this.queue[len(this.queue)-1:]
	defer func() {
		this.temp = append(this.temp, this.queue[0])
		this.queue = this.queue[:0]
		this.queue, this.temp = this.temp, this.queue
	}()
	return this.queue[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue)+len(this.temp) == 0
}
