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

// 迭代
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := ListNode{} // 哨兵节点
	cur := &dummy
	carry := 0 // 进位
	// 循环当前节点
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: carry % 10} // 当前节点的下一个节点
		carry = carry / 10                    // 全局变量， 更新下一个进位
		cur = cur.Next                        //变成下一个节点， 继续便利
	}
	return dummy.Next
}
