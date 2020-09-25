package lbTable4

/**
	算法描述
	给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
 */
 /**
 	声明单链表结构体
  */
type ListNode struct {
	     Val int
	     Next *ListNode
}
func rotateRight(head *ListNode, k int) *ListNode {
	//判断
	if k <= 0 || head == nil {
		return head
	}

	//辅助数组
	var lTable []*ListNode

	//利用循环链表
	p := head
	for p != nil {

		//写入
		if p != nil {
			lTable = append(lTable, p)
		}
		p = p.Next
	}

	//首位相连
	lTable[len(lTable) - 1].Next = lTable[0]

	//根据k 来判断循环链表起始位置
	mod := k % len(lTable)
	begin := len(lTable) - 1
	if mod > 0 {
		begin = len(lTable) - mod
	}

	//断开循环链表
	if begin - 1 >=0 {
		lTable[begin - 1].Next = nil
	} else {
		lTable[len(lTable) - 1].Next = nil
	}


	//返回
	return lTable[begin]

}
