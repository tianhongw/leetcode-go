package a

type MinStack struct {
	items []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)

	if len(this.items) == 1 {
		this.mins = append(this.mins, x)
	} else {
		if this.mins[len(this.items)-2] <= x {
			this.mins = append(this.mins, this.mins[len(this.items)-2])
		} else {
			this.mins = append(this.mins, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.items) == 0 {
		return
	}

	this.items = this.items[0 : len(this.items)-1]
	this.mins = this.mins[0 : len(this.mins)-1]
}

func (this *MinStack) Top() int {
	if len(this.items) == 0 {
		return 0
	}

	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
