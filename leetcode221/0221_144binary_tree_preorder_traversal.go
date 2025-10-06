package leetcode

import "container/list"
import . "lcutils/treenode"

/*
*

	 type Element struct {
	    // 元素保管的值
	    Value interface{}
	    // 内含隐藏或非导出字段
	}

func (l *List) Back() *Element
前序遍历：中左右
压栈顺序：右左中

	*
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = list.New() //栈
	res := []int{}         //结果集
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)     //弹出元素
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：右左中
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
	}
	return res

}
