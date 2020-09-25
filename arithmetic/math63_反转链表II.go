package lbTable5

/**
	算法描述：
	反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

 */
/**
	声明单链表结构体
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//判断参数
	if head == nil || m >=n || head.Next == nil {
		return head
	}

	//声明旋转后的前置
	 pre := head

	 //声明计数
	 count := 1

	 //开始反转
	 p := head

	 //上一个
	 last := head

	 for p != nil && count <= n {
	 	//判断
	 	if count == m - 1 {//m位置的前一个,m==1时 pre = head
	 		pre = p
	 		p = p.Next
	 		count++
	 		continue
		}

	 	//判断是否需要反转
	 	if count == m {
	 		last = p
			p = p.Next
			count++
			continue
		}

	 	//旋转
	 	if count > m {
	 		temp := p
	 		last.Next = p.Next
	 		p.Next = last
	 		last = p
	 		p = temp.Next
			count++
	 		continue

		}

	 	//向下
	 	p = p.Next
	 	count++

	 }

	//拼接
	if pre == head {
		pre.Next = p

	} else {
		pre.Next = last
	}

	 //返回
	 return head

}
/**

 */
