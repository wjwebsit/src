package main

import "fmt"

/**
	算法描述
	给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。

如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。

 

示例 1：

输入：1
输出：true
示例 2：

输入：10
输出：false
示例 3：

输入：16
输出：true
示例 4：

输入：24
输出：false
示例 5：

输入：46
输出：true
 

提示：

1 <= N <= 10^9

 */
func reorderedPowerOf2(N int) bool {
	//判断
	if N == 1 {
		return true
	}

	//获取N的各位数字
	var digit []int
	for N != 0 {
		d := N % 10
		digit = append(digit,d)
		N /= 10
	}

	//返回
	return reorderedPowerOf2HS(digit,0,1,len(digit))
}

func reorderedPowerOf2HS(digit []int,num,cur,n int) bool{
	//判断
	if cur > n {

		return isPowerOf2(num)
	}

	for i := 0; i < len(digit);i ++ {
		//判断
		if cur == 1 && digit[i] == 0 {
			continue
		}

		//当前
		c := digit[i]

		//删除元素
		digit = append(digit[:i],digit[i + 1:]...)

		//dfs
		res := reorderedPowerOf2HS(digit,num * 10 + c,cur + 1,n)

		//判断
		if res == true {
			return true
		}

		//回溯
		tmp := append([]int{c},digit[i:]...)
		digit = append(digit[:i],tmp...)


	}
	//返回
	return false
}



/**
	2的幂满足条件--N & (N - 1) == 0
 */
func isPowerOf2(N int ) bool {
	return N & (N - 1) == 0
}

func main() {
	num := 16
	fmt.Println(reorderedPowerOf2(num))

}