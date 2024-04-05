package leetcode20240405

type MyStack struct {
	inputQueue  []int
	outputQueue []int
}

func Constructor() MyStack {
	return MyStack{
		inputQueue:  make([]int, 0),
		outputQueue: make([]int, 0),
	}

}

func (this *MyStack) Push(x int) {
	this.inputQueue = append(this.inputQueue, x)

}

func (this *MyStack) Pop() int {
	for len(this.inputQueue) != 1 {
		this.outputQueue = append(this.outputQueue, this.inputQueue[0])
		this.inputQueue = this.inputQueue[1:]
	}
	top := this.inputQueue[0]
	this.inputQueue = this.outputQueue
	this.outputQueue = []int{}
	return top

}

func (this *MyStack) Top() int {
	for len(this.inputQueue) != 1 {
		this.outputQueue = append(this.outputQueue, this.inputQueue[0])
		this.inputQueue = this.inputQueue[1:]
	}
	top := this.inputQueue[0]
	this.inputQueue = append(this.outputQueue, this.inputQueue...)
	this.outputQueue = []int{}
	return top
}

func (this *MyStack) Empty() bool {
	if (len(this.inputQueue) == 0) && (len(this.outputQueue) == 0) {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
