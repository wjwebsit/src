package main

import (
	"fmt"
	"sort"
)

/**
算法描述
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
 */
func combinationSum(candidates []int, target int)[][]int{
	//声明返回
	var rtResult [][]int
	if len(candidates) == 0 {
		return rtResult
	}
	sort.Ints(candidates)

	temp := []int{}

	f (target,candidates,temp,0,&rtResult)

	return rtResult
}
/**
	回溯决策
 */
func f (target int, arr []int,temp []int,num int,result *[][]int) {
	//终止条件
	if target < 0 {
		return
	}

	//判断是否为0
	if target == 0 {
		*result = append(*result,temp)
		return
	}

	//当下的决策
	for i := num ; i < len(arr) && arr[i] <= target; i++ {
		//决策
		temp1 := make([]int,len(temp)) //深拷贝
		copy(temp1,temp)
		temp1 = append(temp1,arr[i])
		f (target - arr[i],arr,temp1,i,result)
	}

}

func main() {
	arr := []int{2,3,5}
	fmt.Println(combinationSum(arr,8))
}

