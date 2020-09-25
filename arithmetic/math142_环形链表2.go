package main
/**
	题目描述见 leetCode
 */
func detectCycle(head *ListNode) *ListNode {
	//先判断有没有环---代码和上道题一样

	//判断根
	if head == nil {//head为nil
		return head
	}
	//默认无环
	res := false

	//利用双指针--同一起点
	quick,slow := head,head

	//循环
	for quick != nil && quick.Next != nil{
		//快的跑2步
		quick = quick.Next.Next

		//慢的跑一部
		slow = slow.Next

		//判断是否相等
		if quick == slow {
			res = true
			break
		}

	}

	//如果无环返回nil
	if res == false {
		return nil
	}

	//判断环的长度大小
	length := 1
	quick = quick.Next
	for quick != slow {
		quick = quick.Next
		length++
	}

	//从开头遍历 如果开头往下顺延length后为它本身 说明就是起点
	p := head
	for p != nil {
		//计数
		count := 0
		next := p
		for count < length {
			next = next.Next
			count ++
		}

		//判断
		if next == p {
			break
		}

		p = p.Next

	}

	//返回
	return p
}

