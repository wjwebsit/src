package main

import "fmt"

/**
	算法描述:
	给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5
 */

/**
定义链表
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//尾插法构造单链表 ***
func createTail(L *ListNode,arr []int) *ListNode{
	//构造起始节点和计数器
	start := new(ListNode)
	(*start).Val = arr[0]
	L.Next = start

	count := 1

	for i := 1; i < len(arr);i++ {
		node := new(ListNode)
		(*node).Val = arr[i]
		start.Next = node
		start = start.Next
		count ++
	}

	(*L).Val = count

	return L.Next

}

/**
*@params L 头指针
 */
func read(L *ListNode) {
	p := L
	for p != nil {
		fmt.Print((*p).Val)
		p = p.Next
	}

}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//判断参数
	if head == nil || m >=n || head.Next == nil {
		return head
	}

	//声明旋转后的前置
	pre := head

	//声明计数
	count := 1

	//开始反转
	p := head

	//上一个
	last := head
	tail := head //尾

	for p != nil && count <= n {
		//判断
		if count == m - 1 {//m位置的前一个,m==1时 pre = head
			pre = p
			p = p.Next
			count++
			continue
		}

		//判断是否需要反转
		if count == m {
			last = p
			tail = p
			p = p.Next
			count++
			continue
		}

		//旋转
		if count > m {
			temp := p.Next
			p.Next = last
			last = p
			p = temp
			count++
			continue

		}

		//向下
		p = p.Next
		count++

	}
	tail.Next = p

	//拼接
	if m == 1 {
		return last

	} else {

		pre.Next = last
	}

	//返回
	return head

}

/**
	思路类似于快排序子函数
 */
func partition(head *ListNode, x int) *ListNode {
	//判断
	if head == nil || head.Next == nil {
		return head
	}


	//定义辅助变量
	var top,pre,preLast *ListNode
	firstNode := head

	//循环
	p := head
	for p != nil {
		temp := p.Next
		//判断当前是否小于x
		if p.Val < x {
			//是否为第一次
			if top == nil {
				firstNode = p
				top = p
				pre = p
			} else {
				pre.Next = p.Next
				p.Next = top.Next
				top.Next = p
				top = p
				if pre.Val < x {
					pre = p
				}

			}

		} else {
			pre = p
			if top == nil {
				preLast = p
				//fmt.Println(preLast.Val)
			}
		}
		p = temp

	}

	//判断
	if preLast != nil && top != nil{

		preLast.Next = top.Next
		top.Next = head
	}

	//返回
	return firstNode


}

func main() {
	L := new (ListNode)

	arr := []int{1,2,-1,2,1}

	head := createTail(L,arr)

	p:= partition(head,3)
	read(p)



	//read(head)

}
