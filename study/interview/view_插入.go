package main

/**
	算法描述
	插入。给定两个32位的整数N与M，以及表示比特位置的i与j。编写一种方法，将M插入N，使得M从N的第j位开始，到第i位结束。假定从j位到i位足以容纳M，也即若M = 10 011，那么j和i之间至少可容纳5个位。例如，不可能出现j = 3和i = 2的情况，因为第3位和第2位之间放不下M。

示例1:

 输入：N = 10000000000, M = 10011, i = 2, j = 6
 输出：N = 10001001100
示例2:

 输入： N = 0, M = 11111, i = 0, j = 4
 输出：N = 11111

 */
func insertBits(N int, M int, i int, j int) int {
	//获取N i和j之间的值
	carry := 1
	for k := 1; k <= i ; k ++ {
		carry *= 2
	}

	//获取插值
	sum := 0
	sum2 := 0
	for f := i; f <= j ; f ++ {
		//获取n位的值
		m := (N >> byte(f)) & 1
		n := M >> byte(f - i) & 1

		//获取值
		sum = sum + m * carry
		sum2 = sum2 + n * carry
		carry *= 2
	}


	//返回
	return N + sum2 - sum
}