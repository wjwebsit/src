package main
/**
	算法描述：
	设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
 */
type MinStack struct {
	top1 *node //存储栈top
	top2 *node  //最小栈top

}

type node struct {
	val int
	next *node
}


/** initialize your data structure here. */
func Constructor() MinStack {
	//初始化
	stack := new(MinStack)
	stack.top1 = nil
	stack.top2 = nil

	//返回
	return *stack
}


func (this *MinStack) Push(x int)  {
	//入存储栈
	newNode1 := new(node)
	newNode1.val = x
	newNode2 := new(node)
	newNode2.val = x
	if this.top1 == nil {
		this.top1 = newNode1
		this.top2 = newNode2

	} else {
		newNode1.next = this.top1
		this.top1 = newNode1
		if x <= this.top2.val {
			newNode2.next = this.top2
			this.top2 = newNode2
		}
	}

}


func (this *MinStack) Pop()  {
	if this.top1 != nil {
		//存储出栈
		top := this.top1
		this.top1 = top.next

		//最小栈
		if top.val == this.top2.val {
			this.top2 = this.top2.next
		}

	}
}


func (this *MinStack) Top() int {
	return this.top1.val
}


func (this *MinStack) GetMin() int {
	return  this.top2.val
}
