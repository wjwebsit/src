package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/**
	利用：bit数组可以判断一个整数是否在集合中
 */

type bit struct {
	words []uint64

}

/**
	判断一个元素是否在集合中
 */
func (b *bit) has(x int) bool {
	//获取索引和偏移量
	index , ll := x / 64, x % 64

	//判断
	if len(b.words) <= index {
		return false
	}

	//返回---ll位置是否为1
	return b.words[index] >> ll & 1 == 1

}

/**
	添加元素
 */
func (b *bit) set(x int) bool {
	//获取偏移量
	index,ll := x / 64, x % 64

	//循环
	for len(b.words) <= index {
		b.words = append(b.words,0)
	}

	//ll位置标记为1
	b.words[index] |= 1 << ll

	//返回
	return b.has(x)

}
/**
	删除元素
 */
func (b *bit) del(x int) bool {
	//判断是否存在
	if b.has(x) == false {
		return true
	}

	//获取索引和偏移量
	index,ll := x / 64,x % 64

	//删除
	b.words[index] &= ^(1 << ll)

	//缩容
	b.bizBit()

	//返回结果
	return !b.has(x)
}
/**
	合并bit
 */
func (b *bit) unionBit(y bit) bool {
	//遍历
	for i := 0 ;i < len(y.words) ; i ++ {
		if i < len(b.words) {
			//更新
			b.words[i] |= y.words[i]

		} else {//整行添加
			b.words = append(b.words,y.words[i])
		}
	}

	//返回
	return true
}
/**
	缩容
 */
func (b *bit) bizBit() {
	for len(b.words) > 0 && b.words[len(b.words) - 1] == 0 {
		b.words = b.words[:len(b.words) - 1]
	}

}

/**
	字符串形式打印set
 */
func (b *bit) toString() string {
	//声明buffer
	var buff bytes.Buffer
	buff.WriteByte('{')

	//遍历并检索数字
	for i := 0 ; i < len(b.words) ;i ++ {
		//判断当前是否不存在元素
		if b.words[i] == 0 {
			continue
		}

		//获取元素
		for j := 0; j < 64;j++ {//这里可以优化
			if b.words[i] >> j & 1 == 1 {//有数字
				buff.WriteString(strconv.Itoa(i * 64 + j))
				buff.WriteByte(',')
			}
		}
	}

	//移除最后的','
	buff.Truncate(buff.Len() - 1)
	buff.WriteByte('}')

	//返回
	return buff.String()
}

func main() {
	var b bit
	b.set(12)
	b.set(13)
	b.set(14)
	b.set(16)
	b.set(99999)
	fmt.Println(b.has(13))
	fmt.Println(b.words)
	fmt.Println(b.toString())

	//删除
	b.del(99999)
	fmt.Println(b.words)
	fmt.Println(b.toString())

	var b2 bit
	b2.set(12)
	b2.set(20)
	b2.set(22)
	b2.unionBit(b)
	fmt.Println(b2.toString())
}