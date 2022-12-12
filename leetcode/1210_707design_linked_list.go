package leetcode

type MyLinkedList struct {
	// 如何确定结构体里应设置什么变量？
	// 总的List 包含 有多少个节点，以及头节点是什么
	Size      int
	dummyHead *ListNode
}

func Constructor() MyLinkedList {
	// Constructor 该如何写？如何才能符合题目要求？
	return MyLinkedList{0, &ListNode{}} // 这个虚拟头节点不应该删掉吗？
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	//dummyHead := &ListNode{0, this.head}
	cur := this.dummyHead.Next
	for index != 0 && cur != nil {
		cur = cur.Next
		index--
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= 0 && index <= this.Size {
		cur := this.dummyHead //cur 等于虚拟头节点，插入节点的前驱
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		newNode := &ListNode{val, cur.Next}
		cur.Next = newNode
		this.Size++
	}
	if index < 0 {
		this.AddAtIndex(0, val)
		this.Size++
	}
	if index > this.Size {
		return // 这个return 到哪里了？ 代表结束这个程序吗？
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	if index >= 0 && index < this.Size {
		cur := this.dummyHead //cur 等于虚拟头节点，插入节点的前驱
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		this.Size--
	}
	return
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
