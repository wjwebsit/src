package arithmetic

/**
	算法描述
	将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
	示例：

   输入：1->2->4, 1->3->4
   输出：1->1->2->3->4->4
 */

type ListMathNode struct {
	Val int
	Next *ListMathNode
}
/**
合并有序链表,头插法为逆序 ，采用尾插法
 */
func mergeTwoLists(l1 *ListMathNode, l2 *ListMathNode) *ListMathNode {
	//声明2个变量避免破坏原来的结点
	p1,p2 := l1,l2

	//声明返回值头--尾插法构成新链表
	rtList := new(ListMathNode)
	tempList := rtList

	//循环条件
	for p1 != nil || p2 != nil {
		//第二个链表结束
		for p1 != nil && p2 == nil {
			tempList.Next = p1
			p1 = p1.Next
			tempList = tempList.Next
		}

		//第一个链表结束
		for p1 == nil && p2 != nil {
			tempList.Next = p2
			p2 = p2.Next
			tempList = tempList.Next
		}

		//p1 >= p2 时
		for p1 != nil && p2 != nil && p1.Val >= p2.Val {
			tempList.Next = p2
			tempList = tempList.Next
			p2 = p2.Next

		}

		//p2 > p1
		for p1 != nil && p2 != nil && p2.Val > p1.Val {
			tempList.Next = p1
			p1 = p1.Next
			tempList = tempList.Next
		}

	}

	//返回值
	return rtList.Next
}
