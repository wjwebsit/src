package main
/**
	算法描述:
	插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。
 

示例 1：

输入: 4->2->1->3
输出: 1->2->3->4

 */
type Node struct {
	     Val int
	     Next *Node
}

func insertionSortList(head *Node) *Node {
	//判断指针
	if head == nil {
		return head
	}

	//声明排序号的链表的开始和结束指针--开始为head
	var tail *Node
	tail = head

	//遍历下一次
	next := head.Next

	for next != nil {
		//记录next.Next
		tmp := next.Next

		//判断是否next是否大于tail
		if next.Val >= tail.Val {//直接插入
			tail = next

		} else {
			//从头开始查找
			p := head
			var pre *Node

			for p.Val <= next.Val {
				pre = p
				p = p.Next
			}

			//判断pre
			if pre == nil {
				//插入最前面
				next.Next = head
				head = next

			} else {
				next.Next = pre.Next
				pre.Next = next
			}

		}

		//继续遍历
		next = tmp
	}

	//返回
	return head

}