package main
/**
	算法描述
	实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。

 */
func kthToLast(head *ListNode, k int) int {
	//声明数组
	var list []int

	//获取值
	p := head
	for p != nil {
		list = append(list,p.Val)
		p = p.next
	}

	//返回
	return list[len(list) - k]
}