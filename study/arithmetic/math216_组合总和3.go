package main

import "fmt"

/**
	算法描述
	找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]

 */
func combinationSum3(k int, n int) [][]int {
	//声明返回
	var result [][]int

	//判断
	if k <= 0 || n <= 0 {
		return result
	}

	//采用回溯算法
	combinationSum3Hs(1,n,k,[]int{},&result)

	//返回
	return result
}

func combinationSum3Hs(begin ,target ,k int,tmp []int,result *[][]int)  {

	//判断如果次数用完
	if k == 0 {
		if target == 0 {//表示刚好可以
			//写入结果
			tmp2 := make([]int,len(tmp))
			copy(tmp2,tmp)
			*result = append(*result,tmp2)
			return
		} else {
			return
		}
	}

	//开始

	for num := begin; num <= 9 ; num ++ {
		//判断
		if num * k > target {
			break
		}

		//写入结果
		tmp = append(tmp,num)

		//回溯
		combinationSum3Hs(num + 1,target - num,k - 1,tmp,result)

		//回溯
		tmp = tmp[:len(tmp) - 1]

	}

}
func main() {
	fmt.Println(combinationSum3(3,9))
}