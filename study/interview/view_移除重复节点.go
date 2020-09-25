package main
/**
	算法描述
	编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：

如果不得使用临时缓冲区，该怎么解决？
 */

func removeDuplicateNodes(head *ListNode) *ListNode {
	//利用缓冲区
	hash := make(map[int]bool)
	hash[head.Val] = true

	p := head.Next
	pre := head
	for p != nil {
		if hash[p.Val] == true {
			f := p.Next
			for  f != nil && hash[f.Val] == true {
				f = f.Next
			}

			//更新
			pre.Next = f
			if f == nil {
				break
			}
			hash[f.Val] = true
			p = f.Next
			pre = f

		} else {
			pre = p
			hash[p.Val] = true
			p = p.Next
		}

	}
	pre.Next = nil
	//返回
	return head


}