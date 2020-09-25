package lb

/**
	算法描述
	给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
 */
type ListNode struct {
	     Val int
	     Next *ListNode
	 }

/**
	删除排序链表的重复项
 */
func deleteDuplicates(head *ListNode) *ListNode {
	//判断链表
	if head == nil || head.Next == nil {
		return head
	}

	//定义双指针
	first,next := head,head.Next

	//遍历
	for next != nil {
		fg := false  //交换标识
		//判断相等
		for next != nil && next.Val == first.Val {
			next = next.Next
			fg = true
		}

		//连接
		if fg == true {
			first.Next = next
		}

		//前移
		if next != nil {
			first = first.Next
			next = next.Next
		}
	}

	//返回
	return head
}