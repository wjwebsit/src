package main

import "sync"

/**
	算法描述:
	在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

 */

 /**
 	根据描述需要采用归并排序
  */
func sortList(head *ListNode) *ListNode {
	//判断
	if head == nil {
		return  head
	}

	return mySortList(head)
}

func mySortList(head *ListNode) *ListNode {

	//判断是否只有一个元素
	if head.Next == nil {
		return head
	} else {
		//获取中点 快慢指针
		quick,slow := head.Next,head

		for quick != nil && quick.Next != nil {
			quick = quick.Next.Next
			slow = slow.Next
		}

		quick = slow.Next

		//中断
		slow.Next = nil

		//开始到中心位置链表
		wg := sync.WaitGroup{}
		wg.Add(2)
		first,second := make(chan *ListNode),make(chan *ListNode)
		go func(head *ListNode) {
			first <- mySortList(head)
			close(first)
			wg.Done()
		}(head)

		//中心到结束位置
		go func(quick *ListNode) {
			second <- mySortList(quick)
			close(second)
			wg.Done()
		}(quick)
		wg.Wait()


		//合并
		return mergeSortNode(<-first,<-second)
	}
}
func mergeSortNode(head1,head2 *ListNode) *ListNode {
	//构造合并后的新链表
	var newList *ListNode
	newList = new(ListNode)
	tail := newList

	//合并2个有序链表
	p1,p2 :=  head1,head2

	for p1 != nil && p2 != nil {
		//p1 >= p2
		for p2 != nil && p1.Val >= p2.Val {
			tail.Next = p2
			tail = p2
			p2 = p2.Next
		}
		//p2.Val > p1.Val 可能p2为空零
		for p1 != nil && p2 != nil && p2.Val > p1.Val {
			tail.Next = p1
			tail = p1
			p1 = p1.Next
		}
	}

	//拼接最后
	if p1 != nil {
		tail.Next = p1
	}

	if p2 != nil {
		tail.Next = p2
	}

	//返回
	return newList.Next

}