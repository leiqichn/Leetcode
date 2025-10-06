package _0251006_dfs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	return dfs(nil, head)
}

func dfs(pre *ListNode, cur *ListNode) *ListNode {

	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre

	return dfs(cur, next)

}
