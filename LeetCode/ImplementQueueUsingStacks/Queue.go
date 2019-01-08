package ImplementQueueUsingStacks

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		input:  make([]int, 0, 5),
		output: make([]int, 0, 5),
	}
}

type MyQueue struct {
	input, output []int
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	defer func() {
		this.output = this.output[:len(this.output)-1]
	}()
	return this.Peek()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	if len(this.output) == 0 {
		for i := len(this.input) - 1; i >= 0; i-- {
			this.output = append(this.output, this.input[i])
		}
		this.input = this.input[:0]
	}
	return this.output[len(this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.input)+len(this.output) == 0
}
