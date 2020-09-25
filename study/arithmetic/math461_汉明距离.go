package main

/**
	算法描述:
	两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。

给出两个整数 x 和 y，计算它们之间的汉明距离。

注意：
0 ≤ x, y < 231.

示例:

输入: x = 1, y = 4

输出: 2

解释:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

上面的箭头指出了对应二进制位不同的位置。
 */
func hammingDistance(x int, y int) int {
	//判断
	if x == y {
		return 0
	}

	//做异或操作
	z := x ^ y

	//统计z中1 的数目
	sum := 0
	for z > 0 {
		if z & 1 == 1 {
			sum ++
		}
		z >>= 1
	}

	//返回
	return sum


}