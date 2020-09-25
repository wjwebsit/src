package main

/**
	算法描述：（略）
 */
//定义结构体
type Node struct {
	 Val int
	 Next *Node
	 Random *Node
}


func copyRandomList(head *Node) *Node {
	//判断
	if head == nil {
		return nil
	}

	//声明hash
	var hash = make(map[*Node]*Node)

	p := head
	var newHead *Node
	var pre *Node

 	for p != nil {
 		//判断缓存中是否存在
 		if _,ok :=  hash[p];!ok {
 			hash[p] = new (Node)
 			hash[p].Val = p.Val
		}
 		//是否为第一次进入
 		if pre != nil {
			//连接
			pre.Next = hash[p]

			//指向新结点---赋值随机指针
			pre = pre.Next

		} else {
			pre = hash[p]

			//深拷贝后的新头
			newHead = pre
		}

 		//复制random
 		if p.Random != nil {
 			//判断是否在缓存
 			if _,ok := hash[p.Random];!ok {
				hash[p.Random] = new (Node)
				hash[p.Random].Val = p.Random.Val
			}
 			pre.Random = hash[p.Random]
		}

 		//继续
 		p = p.Next

	}

	//返回
	return newHead

}
