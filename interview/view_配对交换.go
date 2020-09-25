package main
/**
	算法描述
	配对交换。编写程序，交换某个整数的奇数位和偶数位，尽量使用较少的指令（也就是说，位0与位1交换，位2与位3交换，以此类推）。

示例1:

 输入：num = 2（或者0b10）
 输出 1 (或者 0b01)
示例2:

 输入：num = 3
 输出：3
提示:

num的范围在[0, 2^30 - 1]之间，不会发生整数溢出。

 */
func exchangeBits(num int) int {
	//判断
	if num == 0 {
		return 0
	}
	//进位
	carry := 1

	//最终交换完的代价
	sum := 0

	var i byte
	//记录
	for  i = 0; num >> i & 1 > 0 || (num >> i + 1) & 1 > 0; i += 2 {
		//计算代价
		sum += exchangeValue(num >> (i + 1) & 1,num >> i,carry)


		//carry
		carry >>= 2

	}

	//返回
	return num + sum

}
/**
	交换之后的代价，a为高位，b为低位 carry 为当前的地位二进制
*/
func exchangeValue(a,b int,carry int) int {
	if a == b {
		return 0
	}

	if a == 1 && b == 0 {
		return  -1 * carry
	}
	return carry
}