package main
/**
	算法描述
	给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。

示例 1：

输入: num = 1775(110111011112)
输出: 8
示例 2：

输入: num = 7(01112)
输出: 4
 */
func reverseBits(num int) int {
	//利用双指针
	var i  = 0

	//上一次结束
	var last = -1
	var lastLength = 0

	//定义返回值
	var max = 0

	//循环获取第i位
	for i <= 32 {

		//获取位数
		digit := num >> byte(i) & 1

		//判断第一个1
		if digit == 1 {
			//当前开始
			s := i
			e := i + 1
			for e <= 32 && (num >> byte(e) & 1) == 1 {
				e ++
			}

			//当前长度
			length := e - s

			//判断上一次结束是否相隔1
			if last != -1 && last == s - 1 {
				length += lastLength
			}

			//比较
			if length + 1 > max {
				max = length + 1
			}

			//重置
			last = e
			lastLength = e - s

			//继续 --此时的e不为1
			i = e + 1

		} else {
			i ++
		}

	}

	//判断
	if max > 32 {//全是1
		max = 32
	}
	if max == 0 {//全是0
		max = 1
	}

	return max

}