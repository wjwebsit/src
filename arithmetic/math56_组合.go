package main

import "fmt"

/**
	算法描述
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
 */
 /**
 	回溯算法
  */
func combine(n int, k int) [][]int {
	//声明返回结果
	var result [][]int

	//判断
	if n < k {
		return result
	}

	//初始化
	var tmp []int
	var choice []int
	for i := 1;i <= n; i++ {
		choice = append(choice,i)
	}

	f9(k,choice,0,0,tmp,&result)

	//返回结果
	return result

}
/**
	回溯
 */
func f9(k int,choice []int,height int,start int,tmp []int,result *[][]int) {
	//判断
	if height == k {
		//写入结果
		*result = append(*result,tmp)
		return
	}

	//判断越界
	if height > k {
		return
	}

	//遍历
	for start < len(choice) && height + 1 <= k {
		//深拷贝
		tmp1 := make([]int,len(tmp))
		copy(tmp1,tmp)

		//写入当前
		tmp1 = append(tmp1, choice[start])

		//继续决策
		f9(k,choice,height + 1,start + 1,tmp1,result)

		start ++
	}

}

func main() {
	fmt.Println(combine(4,5))
}
