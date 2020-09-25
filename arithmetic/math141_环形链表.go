package main
/**
	算法描述

 */

func hasCycle(head *ListNode) bool {
	//定义返回值
	res := false

	//判断根
	if head == nil {//head为nil
		return res
	}

	//利用双指针--同一起点
	quick,slow := head,head

	//循环
	for quick != nil && quick.Next != nil {
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

	//返回
	return res

}