package lb1

import "math"

/**
	算法描述
	给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
 */

 /**
 	定义链表的结构体
  */
type ListNode struct {
	Val int
	Next *ListNode
}
/**
	删除数组中的重复项
 */
func deleteDuplicates(head *ListNode) *ListNode {
	//判断
	if head == nil || head.Next == nil {
		//返回原来的链表
		return head
	}

	//定义栈
	stack := [] *ListNode{}

	//head 入栈
	stack = append(stack,head)

	//从下一个开始
	next := head.Next

	//pre --用于空栈时判断1-1-1
	pre := head

	for next != nil {
		//判断是否栈空
		if len(stack) == 0 && pre.Val != next.Val {//入栈
			stack = append(stack,next)

		} else if len(stack) > 0 && stack[len(stack) - 1].Val == next.Val {//和栈顶元素相等--弹出
			stack = stack[:len(stack) - 1]
			//前一个
			pre = next

			//判断是否为
			if len(stack) > 0 {
				stack[len(stack) - 1] .Next = nil
			}

		} else if pre.Val != next.Val{//入栈
			stack[len(stack) - 1].Next = next
			stack = append(stack,next)
		}

		next = next.Next
	}


	//返回
	if len(stack) == 0 {
		return nil
	} else {
		return stack[0]
	}

}