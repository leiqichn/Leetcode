package leetcode141

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 一个走一步，一个走两步，如果不会相遇则不是环形
func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for ; slow.Next != nil; slow = slow.Next {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			break // 如果是slow 在遍历，需要增加这个，否则slow 会追上来
		}
		if slow == fast {
			return true
		}
	}
	return false
}

// 方法2 优化，使用类似whiel 来不断遍历，直到fast 为空或者fast.Next 为空就停止，fast.Next.Next是可以的
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
