package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 新建 虚拟头节点node
	dummyNode := &ListNode{0, head}
	// 遍历 node ，if node.next == val ; node.next= node.next.next；否则 node= node.next
	// 终止条件 for node.next != nil
	node := dummyNode
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next // node.Next.Next 是可以为空的，例如 node= 5,node.next=6, node.next.next =nil
		} else { // 注意else 的话，需要将node的赋值为下一个next，这样node的遍历才可以完成
			node = node.Next
		}
	}
	return dummyNode.Next
}
