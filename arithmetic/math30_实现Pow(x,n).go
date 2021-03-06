package main

import "fmt"

/**
	算法描述
	实现  pow（x，n）  ，即计算x的n次幂函数。

使用折半计算，每次把n缩小一半，这样n最终会缩小到0，
任何数的0次方都为1，这时候我们再往回乘，
如果此时n是偶数，直接把上次递归得到的值算个平方返回即可，
如果是奇数，则还需要乘上个x的值。
还有一点需要引起我们的注意的是n有可能为负数，
对于n是负数的情况，我们可以先用其绝对值计算出一个结果再取其倒数即可。
我们让i初始化为n，然后看i是否是2的倍数，是的话x乘以自己，否则res乘以x，i每次循环缩小一半，直到为0停止循环。最后看n的正负，如果为负，返回其倒数。

示例1：

输入： 2.00000,10
 输出： 1024.00000
示例2：

输入： 2.10000,3
 输出： 9.26100
示例3：

输入： 2.00000，-2
 输出： 0.25000
 解释： 2 -2 = 1/2 2 = 1/4 = 0.25
 */
func myPow(x float64, n int) float64 {
	//特殊情况
	if n == 0 || x == 1.0{
		return 1
	}

	//判断 n
	if n < 0 {
		x = 1/x
		n = -1 * n //将幂变为正整数
	}

	//定义返回值
	rtResult := 1.0

	for j := n; j != 0; {
		if j % 2 != 0 {
			rtResult *= x
		}
		x *= x

		j >>= 1
	}

	//返回结果
	return rtResult

}

func  main ()  {
	fmt.Println(myPow(3.0,3))
}