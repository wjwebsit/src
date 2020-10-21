package main
/**
	算法描述
	在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 个数字。

注意:
n 是正数且在32位整数范围内 ( n < 231)。

示例 1:

输入:
3

输出:
3
示例 2:

输入:
11

输出:
0

说明:
第11个数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是0，它是10的一部分。
 */
func findNthDigit(n int) int {
	//判断特殊情况
	if n <= 10 {
		return n - 1
	}

	//计算所在几位数区间
	digit,ll := 2,10
	index := 10
	var pre int
	for index < n {
		//记录上次位数的结束
		pre = index

		//当前加上此位数
		index = index + digit * ll * 9

		//准备下一次
		digit ++

		//下一次的每一位的数字个数
		ll *= 10
	}

	//当前位数
	digit -= 1
	ll /= 10

	//获取数字
	n -= pre
	num := 0
	for n > 0 && digit > 1{
		i := 0
		for i < 10 && n > 0{
			n -= digit * ll
			i ++
		}

		//获取位数
		i --
		num += i * ll

		//重新赋值
		n = n - i * ll * digit
		digit -= 1
		ll /= 10
	}

	//最后一位
	if n > 0 {
		num += n - 1
	}




}