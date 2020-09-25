package main
/**
	算法描述
	给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。

 

示例：

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：假设这些数位是正向存放的，请再做一遍。

示例：

输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//声明返回值
	var result,tail *ListNode

	//采用双指针
	var p1,p2 = l1,l2

	//声明进位
	carry := 0

	for p1 != nil && p2 != nil {
		//计算和
		sum := p1.Val + p2.Val + carry

		//进位
		carry = sum / 10

		//当前位值
		cur := sum % 10

		//构造新节点
		newNode := new(ListNode)
		newNode.Val = cur

		if result == nil {
			result = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = tail.Next
		}

		p1 = p1.Next
		p2 = p2.Next
	}


	//判断最后
	if p1 != nil {
		for p1 != nil {
			//求和
			sum := p1.Val + carry
			carry = sum / 10
			cur := sum % 10

			//构造新节点
			newNode := new(ListNode)
			newNode.Val = cur
			if result == nil {
				result = newNode
				tail = newNode
			} else {
				tail.Next = newNode
				tail = tail.Next
			}

			p1 = p1.Next
		}
	}

	if p2 != nil {
		for p2 != nil {
			//求和
			sum := p2.Val + carry
			carry = sum / 10
			cur := sum % 10

			//构造新节点
			newNode := new(ListNode)
			newNode.Val = cur
			if result == nil {
				result = newNode
				tail = newNode
			} else {
				tail.Next = newNode
				tail = tail.Next
			}

			p2 = p2.Next
		}

	}

	//最后的进位
	if carry == 1 {
		newNode := new(ListNode)
		newNode.Val = 1
		tail.Next = newNode
		tail = tail.Next
	}

	//返回
	return result

}