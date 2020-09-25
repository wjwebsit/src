package main

import "fmt"

/**
	算法描述
	运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

进阶:

你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:

LRUCache cache = new LRUCache( 2  缓存容量  );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
 */

type LRUCache struct {
	maxCap int //最大容量
	cap int //当前容量
	data map[int]int //数据存储
	head *BinaryLinkList //当前最不活跃
	tail *BinaryLinkList //当前最活跃度
	degree map[int]*BinaryLinkList //活跃度记录散列 用于修改活跃度
}
/**
双向链表来记录活跃度
 */
type  BinaryLinkList struct {
	 pre ,next *BinaryLinkList
	 val int
}

/**
	初始化
 */
func Constructor1(capacity int) LRUCache {
	var LRU = new(LRUCache)
	LRU.maxCap = capacity
	LRU.cap = 0
	LRU.data = make(map[int]int,capacity)
	LRU.head = nil
	LRU.tail = nil
	LRU.degree = make(map[int]*BinaryLinkList,capacity)

	//返回return
	return *LRU
}


func (this *LRUCache) Get(key int) int {
	//判断是否存在
	if _,ok := this.data[key];ok {
		//修改活跃度
		if  this.cap == 1 {//不用修改
			return this.data[key]

		} else {
			//获取链表元素
			node := this.degree[key]

			//如果就是末尾也不用修改
			if node == this.tail {
				return this.data[key]
			}

			//添加到末尾--并删除原来的
			if node == this.head {//如果为第一个元素 则head直接指向下一个
				this.head = node.next
				node.next.pre = nil

			} else {
				//删除
				node.pre .next = node.next
				node.next.pre = node.pre
			}

			//追加到末尾
			this.tail .next = node
			node.next = nil
			node.pre = this.tail
			this.tail = node

			//返回
			return this.data[key]

		}


	} else {
		//不存在返回 -1
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	//判断key是否存在
	if _,ok := this.data[key];ok  {//key存在
		//修改
		this.data[key] = value

		//修改活跃度
		this.Get(key)

		return
	}
	//构造双向链表
	newNode := new(BinaryLinkList)
	newNode.val = key

	//判断是否容量满了
	if this.maxCap > this.cap {//未满--
		//写入
		this.data[key] = value

		//当前的容量+1
		this.cap ++

		//判断是否为第一个元素
		if this.cap == 1 {
			this.tail = newNode
			this.head = newNode

		} else {
			this.tail.next = newNode
			newNode.pre = this.tail
			this.tail = newNode
		}

		//写入散列
		this.degree[key] = newNode

	} else {//删除头
		//获取head值
		headVal := this.head.val

		//判断最大容量
		if this.maxCap == 1 {
			this.tail = newNode
			this.head = newNode

		} else {
			//修改head
			this.head.next.pre = nil
			this.head = this.head.next

			//添加新节点
			this.tail.next = newNode
			newNode.pre = this.tail
			this.tail = newNode

		}

		//删除最不活跃
		delete(this.data,headVal)
		delete(this.degree,headVal)

		//写入key
		this.data[key] = value
		this.degree[key] = newNode

	}
}
func main() {
	cache := Constructor1(2)
	cache.Put(1, 1);

	cache.Put(2, 2);
	fmt.Println(cache.Get(1));       // 返回  1
	cache.Put(3, 3);    // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2));       // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1));       // 返回 -1 (未找到)
	fmt.Println(cache.Get(3));       // 返回  3
	fmt.Println(cache.Get(4));       // 返回  4
}