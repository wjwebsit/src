package main
/**
	算法描述
	合并  ķ 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度
	示例：

输入：
[
  1-> 4-> 5，
  1-> 3-> 4，
  2-> 6
]
输出： 1-> 1-> 2-> 3-> 4-> 4-> 5-> 6
 */

 /**
 	合并逻辑
  */
func mergeKLists(lists []*ListMathNode) *ListMathNode {
	//判断长度
	if len(lists) == 0 {
		return nil
	}
	//判断长度
	if len(lists) == 1 {
		return lists[0]
	}
	//采用分治策略
	for len(lists) != 1 {
		//合并前2项
		lists = append(lists,mergeTwoList(lists[0],lists[1]))

		//删除前2项
		lists = lists[2:]
	}

	//返回
	return lists[0]



}

type ListMathNode struct {
	Val int
	Next *ListMathNode
}
/**
合并有序链表,头插法为逆序 ，采用尾插法
 */
func mergeTwoList(l1 *ListMathNode, l2 *ListMathNode) *ListMathNode {
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