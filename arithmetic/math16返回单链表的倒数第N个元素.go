package xuxia

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}
/**
	遍历
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//定义遍历过的指针数组
	var nodePtrArr [] *ListNode
	nodePtrArr = append(nodePtrArr,head)

	//定义计数器
	count := 0

	//从头结点开始 第一个元素
	p := (*head).Next
	for  p != nil {
		//计数器+1
		count ++

		//存放元素
		nodePtrArr = append(nodePtrArr,p)

		//继续走
		p = (*p).Next
	}

	//判断总长度和n
	if count < n {
		return head.Next
	}

	//找到倒数n个元素之前的结点,由于第一个节点是头，count + 1是 指针数组的真实长度
	node := nodePtrArr[count - n]

	//删除结点
	node.Next = node.Next.Next

	//返回
	return head




}
