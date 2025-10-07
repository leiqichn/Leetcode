package _0251007

type ListNode struct {
	Val  int
	Next *ListNode
}

// l1 和 l2 为当前遍历的节点，carry 为进位
func addTwo(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 { // 递归边界
		return nil
	}

	s := carry
	if l1 != nil {
		s += l1.Val // 累加进位与节点值
		l1 = l1.Next
	}
	if l2 != nil {
		s += l2.Val
		l2 = l2.Next
	}

	// s 除以 10 的余数为当前节点值，商为进位
	return &ListNode{s % 10, addTwo(l1, l2, s/10)}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}
