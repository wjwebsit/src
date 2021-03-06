package main

import "fmt"

/**
	快速幂：快速幂算法——可迅速求出a^b。其主要理论依据如下：

        1，当b为偶数时，a^b可以转为a^2的b/2次方。

        2，当b为奇数时，a^b可以转为a^2的b/2次方，再乘以a。

快速幂的目的就是做到快速求幂，假设我们要求a^b,按照朴素算法就是把a连乘b次，这样一来时间复杂度是O(b)也即是O(n)级别，快速幂能做到O(logn)，快了好多好多。它的原理如下：
　　假设我们要求a^b，那么其实b是可以拆成二进制的，该二进制数第i位的权为2^(i-1)，例如当b==11时
　　 a^11=a(2^0+2^1+2^3)
　　11的二进制是1011，11 = 2³×1 + 2²×0 + 2¹×1 + 2º×1，因此，我们将a¹¹转化为算 a2^0*a2^1*a2^3，也就是a1*a2*a8 ，看出来快的多了吧原来算11次，现在算三次，但是这三项貌似不好求的样子….不急，下面会有详细解释。 　　由于是二进制，很自然地想到用位运算这个强大的工具：&和>> &运算通常用于二进制取位操作，例如一个数 & 1 的结果就是取二进制的最末位。还可以判断奇偶x&1==0为偶，x&1==1为奇。 >>运算比较单纯,二进制去掉最后一位

 */
func  quickPower(a ,b int) int{
	s := 1
	for b > 0 {
		if b & 1 != 0 { //x&1==0为偶，x&1==1为奇
			s *= a
		}
		a = a * a
		b = b >> 1 //右移1位
	}
	return s
}

/**
	利用反复平方法求 a^b % n
    a^2c mod n=(a^c)^2 mod n 为偶数

    a^(2c+1) mod n =a*(a^c)^2 mod n //为奇数
 */
 func repeatPowerMod(a ,b ,n int) int {
 	ans := 1
 	for b != 0 {
 		if b & 1 == 1 { //表示为奇数
 			ans = ans * a % n
		}
 		a = a * a % n
 		b >>= 1

	}
 	return ans
 }

func main() {
	//测试快速幂
	fmt.Println(quickPower(10,2))

	//测试快速幂%n
	fmt.Println(repeatPowerMod(5,1,24))
}