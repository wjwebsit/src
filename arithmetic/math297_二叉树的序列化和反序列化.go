package xuxiao

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	算法描述:
	序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例: 

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	//判断
	if root == nil {
		return "[]"
	}

	//声明字符串数组
	var strArr []string

	//声明队列
	var queue []*TreeNode
	queue = append(queue,root)

	//采用层序遍历
	for len(queue) > 0 {
		//出队
		cur := queue[0]
		queue = queue[1:]

		//判断是否为nil
		if cur == nil {
			strArr = append(strArr,"null")
			continue
		}

		//当前写入
		strArr  = append(strArr,strconv.Itoa(cur.Val))

		//写入子
		queue = append(queue,cur.Left)
		queue = append(queue,cur.Right)


	}

	//返回
	return "[" + strings.Join(strArr,",") + "]"


}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	//判断
	if data == "[]" {
		return nil
	}

	//转化成数组
	var trees []string
	str := string(data[1:len(data) - 1])
	trees = strings.Split(str,",")

	//生明父
	var parents []*TreeNode

	//采用完全二叉树
	for i := 0; i < len(trees);i++ {
		//判断
		if trees[i] == "null" {//不做处理
			continue
		}

		//获取值
		num,_ := strconv.Atoi(trees[i])

		//构造结点
		node := new(TreeNode)
		node.Val = num

		//判断
		if len(parents) == 0 {
			parents = append(parents,node)

		} else {
			//寻找双亲
			k := (i - 1) / 2

			//判断左孩子还是右孩子
			if 2 * k + 1 == i {
				parents[k].Left = node

			} else {
				parents[k].Right = node
			}

			parents = append(parents,node)
		}

	}

	//返回
	return parents[0]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */