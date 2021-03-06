package main

import "fmt"

/**
	算法描述:
	颠倒给定的 32 位无符号整数的二进制位。

 

示例 1：

输入: 00000010100101000001111010011100
输出: 00111001011110000010100101000000
解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
      因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
示例 2：

输入：11111111111111111111111111111101
输出：10111111111111111111111111111111
解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
      因此返回 3221225471 其二进制表示形式为 10101111110010110010011101101001。

 */
func reverseBits(num uint32) uint32 {
	//声明返回
	var rtNum uint32
	rtNum = 0


	//二进制转换
	var count = 0
	for num > 0 {
		rtNum = rtNum * 2 + num % 2
		num = num / 2
		count ++
	}
	for 32 > count {//前面的0
		rtNum *= 2
		count ++
	}

	//返回
	return  rtNum
}
func main()  {
	fmt.Println(reverseBits(4294967293))
}