package main

import (
	"fmt"
	"strconv"
)

/**
	算法描述
	给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
 */
 /**
 	利用回溯求解
  */
func getPermutation(n int, k int) string {
	//定义变量
	var choice []int
	var result string
	var  tmp string
	isEnd := false
	sum := 0

	//初始化决策
	for i := 1; i <= n; i ++ {
		choice = append(choice,i)
	}


	//决策
	if n > 0 && k >0 {
		f13(choice, len(choice), 0, k, tmp, &sum, &isEnd, &result)
	}

	//返回
	return result
}

/**
	@param choice []int 当前决策
	@param length int 最大深度
	@param height int 递归深度
	@param k int 第k个数
	@param tmp string 当前字符
	@param sum *int 当前序列位置
	@param isEnd *bool 是否结束
	@param result *string 结果
 */
func f13(choice []int,length int,height int,k int,tmp string,sum *int,isEnd *bool,result *string){
	//判断
	if *isEnd == true {
		return
	}

	//判断
	if height == length {
		*sum ++
		if *sum == k {
			fmt.Println(tmp)
			*result = tmp
			*isEnd = true
		}
		return
	}

	//抉择
	for i := 0; i < len(choice) && *isEnd == false; i++ {
		//深拷贝
		choice1 := make([]int,len(choice))
		copy(choice1,choice)

		//记录
		temp := choice[i]

		//删除当前
		if len(choice) > 1 {
			choice1 = append(choice1[:i], choice1[i+1:]...)

		} else {
			choice1 = []int{}
		}

		f13(choice1,length,height + 1,k,tmp + strconv.Itoa(temp),sum,isEnd,result)
	}
}

func main() {
	fmt.Println(getPermutation(20,80))
}