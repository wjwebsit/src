package lTbale2


/**
	算法描述 :
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */

//结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义返回值
	result := new(ListNode)
	tail := result


	//声明进位
	carry := 0

	//遍历末尾---高位
	add1 := l1
	add2 := l2

	//逐位相加
	for add1 != nil || add2 != nil {
		//新的结点
		newNode := new(ListNode)

		//判断情况1 ---add1到尽头
		if add1 == nil && add2 != nil {
			if (add2.Val + carry) >= 10 {
				newNode.Val = (add2.Val + carry) % 10
				carry = 1
			} else {
				newNode.Val = add2.Val + carry
				carry = 0
			}

			//尾插法
			tail.Next = newNode
			tail = tail.Next

			//往下
			add2 = add2.Next
			continue
		}
		//判断情况1 ---add2到尽头
		if add2 == nil && add1 != nil {
			if (add1.Val + carry) >= 10 {
				newNode.Val = (add1.Val + carry) % 10
				carry = 1
			} else {
				newNode.Val = add1.Val + carry
				carry = 0
			}

			//尾插法
			tail.Next = newNode
			tail = tail.Next

			//往下
			add1 = add1.Next
			continue
		}

		//都未
		//判断情况1 ---add2到尽头
		if add2 != nil && add1 != nil {
			if (add1.Val + carry + add2.Val) >= 10 {
				newNode.Val = (add1.Val + carry + add2.Val) % 10
				carry = 1
			} else {
				newNode.Val = add1.Val + carry + add2.Val
				carry = 0
			}


			//尾插法
			tail.Next = newNode
			tail = tail.Next

			//往下
			add1 = add1.Next
			add2 = add2.Next

		}
	}

	//判断进位
	if carry == 1 {
		newNode := new(ListNode)
		newNode.Val = 1

		//尾插法
		tail.Next = newNode
		tail = tail.Next
	}

	//返回
	return result.Next

}
