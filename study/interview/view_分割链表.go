package main
/**
	算法描述
	编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。

示例:

输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8
 */
func partition(head *ListNode, x int) *ListNode {
	//判断
	if head == nil {
		return head
	}

	//声明数组
	var list []*ListNode
	p := head
	for p != nil {
		list = append(list,p)
		p = p.Next
	}

	//利用双指针
	s,e := 0, len(list) - 1
	for s < e {
		//从前边找到一个比x大的
		for s < e && list[s].Val < x {
			s ++
		}

		//从后边找到一个比x小的
		for s < e && list[e].Val >= x {
			e --
		}

		//交换
		list[s].Val,list[e].Val = list[e].Val,list[s].Val

		//继续
		s ++
		e --
	}

	return head

}