package leetcode20240405

type MyQueue struct {
	inputStack  []int
	outputStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inputStack:  make([]int, 0),
		outputStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.inputStack = append(this.inputStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outputStack) != 0 {

		topTmp := this.outputStack[len(this.outputStack)-1] //
		this.outputStack = this.outputStack[:len(this.outputStack)-1]
		return topTmp
	} else {
		for len(this.inputStack) != 0 {
			this.outputStack = append(this.outputStack, this.inputStack[len(this.inputStack)-1])
			this.inputStack = this.inputStack[:len(this.inputStack)-1]
		}
	}

	top := this.outputStack[len(this.outputStack)-1] //
	this.outputStack = this.outputStack[:len(this.outputStack)-1]
	return top
}

func (this *MyQueue) Peek() int {
	if len(this.outputStack) != 0 {

		topTmp := this.outputStack[len(this.outputStack)-1] //
		return topTmp
	} else {
		for len(this.inputStack) != 0 {
			this.outputStack = append(this.outputStack, this.inputStack[len(this.inputStack)-1])
			this.inputStack = this.inputStack[:len(this.inputStack)-1]
		}
	}

	top := this.outputStack[len(this.outputStack)-1] //
	return top
}

func (this *MyQueue) Empty() bool {
	if len(this.inputStack) == 0 && len(this.outputStack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
