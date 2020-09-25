package main
/**
	算法描述
	给定一个迭代器类的接口，接口包含两个方法： next() 和 hasNext()。设计并实现一个支持 peek() 操作的顶端迭代器 -- 其本质就是把原本应由 next() 方法返回的元素 peek() 出来。

示例:

假设迭代器被初始化为列表 [1,2,3]。

调用 next() 返回 1，得到列表中的第一个元素。
现在调用 peek() 返回 2，下一个元素。在此之后调用 next() 仍然返回 2。
最后一次调用 next() 返回 3，末尾元素。在此之后调用 hasNext() 应该返回 false。
进阶：你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？

 */

 /**

  */
type PeekingIterator struct {
	 tail,head *queueNode
}

type queueNode struct {
	data int
	next *queueNode
}



func Constructor6(iter *Iterator) *PeekingIterator {
	//构造
	queue := new(PeekingIterator)
	queue.head = nil
	queue.tail = nil

	//判断
	for iter.hasNext() {
		node := new(queueNode)
		node.data = iter.next()

		//从尾部写入队列
		if queue.tail == nil {
			queue.head = node
			queue.tail = node

		} else {
			queue.tail.next = node
			queue.tail = node
		}
	}

	//返回
	return queue
}

func (this *PeekingIterator) hasNext() bool {
	return this.head != nil
}

func (this *PeekingIterator) next() int {
	rt := this.head.data
	this.head = this.head.next
	return  rt
}

func (this *PeekingIterator) peek() int {
	return this.head.data
}