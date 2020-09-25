package main

type ListMathNode1 struct {
	Val int
	Next *ListMathNode1
}
/**
	交换
 */
func swapPairs(head *ListMathNode1) *ListMathNode1 {
	//判断
	if head == nil {
		return nil
	}

	//获取第一个
	first := head
	if first.Next == nil {
		return first
	}

	//获取第二个
	second := first.Next

	//head 起始指向second 因为第一次交换后此为头
	head = second

	//记录上一轮指向新的second
	pre := first

	//执行交换
	for second != nil {

		//交换
		first.Next = second.Next
		second.Next = first

		//更新下一轮
		first = first.Next
		if first == nil {
			pre.Next = first
			break
		}
		second = first.Next

		//交换完前轮的需要指向second [1,2,3]
		if second != nil {
			pre.Next = second
			pre = first
		}

	}

	//返回
	return head

}
