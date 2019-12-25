package a

type MyQueue struct {
	stack     []int
	helpStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.helpStack = nil
	for i := 0; i < len(this.stack); i++ {
		this.helpStack = append(this.helpStack, this.stack[i])
	}
	this.helpStack = append(this.helpStack, x)

	this.stack = nil
	for i := 0; i < len(this.helpStack); i++ {
		this.stack = append(this.stack, this.helpStack[i])
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	x := this.stack[0]
	this.stack = this.stack[1:]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
