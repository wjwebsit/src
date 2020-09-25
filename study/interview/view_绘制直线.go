package main

import (
	"fmt"
	"math"
)

/**
	绘制直线。有个单色屏幕存储在一个一维数组中，使得32个连续像素可以存放在一个 int 里。屏幕宽度为w，且w可被32整除（即一个 int 不会分布在两行上），屏幕高度可由数组长度及屏幕宽度推算得出。请实现一个函数，绘制从点(x1, y)到点(x2, y)的水平线。

给出数组的长度 length，宽度 w（以比特为单位）、直线开始位置 x1（比特为单位）、直线结束位置 x2（比特为单位）、直线所在行数 y。返回绘制过后的数组。

示例1:

 输入：length = 1, w = 32, x1 = 30, x2 = 31, y = 0
 输出：[3]
 说明：在第0行的第30位到第31为画一条直线，屏幕表示为[0b000000000000000000000000000000011]
示例2:

 输入：length = 3, w = 96, x1 = 0, x2 = 95, y = 0
 输出：[-1, -1, -1]
 */

func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	//声明绘画后的数组
	var result = make([]int,length) //并赋予默认值0

	//判断
	if length == 0 {
		return result
	}

	//定义起始下标
	s := 0


	//判断行数
	if y > 0 {
		s += y * (w / 32)
	}

	//x1和x2 在同一行
	var i = 1  //寻找像素点开始
	for i * 32 < x1 {
		i ++
	}

	if x2 >= i * 32 {
		//当前绘制一部分
		result[s + i - 1] = getBinaryValue(x1 - (i - 1) * 32, 31)

		//剩余绘制
		i ++
		for i *32 <= x2 {
			//直接为-1
			result[s + i - 1] = -1
			i++
		}

		//最后在绘制
		if  i * 32 - x2 < 32 {
			result[s + i - 1] = getBinaryValue(0, x2 - (i - 1) * 32)
		}


	} else {//就在当前i绘制完
		result[s + i - 1] = getBinaryValue(x1 - (i - 1) * 32,x2 - (i - 1) * 32)

	}
	//返回return
	return result
}

func  getBinaryValue(a,b int) int {
	//求解
	sum := 0
	carry := 1
	for i := 1; i < 32 - b; i ++ {
		carry <<= 1
	}
	if a == 0 {
		//计算
		for b >= a+1 {
			sum += carry
			carry <<= 1
			b--
		}

	} else {
		//计算
		for b >= a {
			sum += carry
			carry <<= 1
			b--
		}
	}


	//32位a和b之间为1
	if a != 0 {
		return  sum

	} else {//表示为负数的补码
		return sum - math.MaxInt32 - 1
	}

}
func main() {
	fmt.Println(getBinaryValue(0,9))
}