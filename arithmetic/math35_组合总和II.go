package main

import (
	"fmt"
	"sort"
)

/**
	算法描述：
	给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
 */
func combinationSum2(candidates []int, target int) [][]int {
	//定义返回值
	result := [][]int{}

	//判断
	if len(candidates) == 0 {
		return result
	}

	//排序
	sort.Ints(candidates)


	temp := []int{}
	f2(target,candidates,temp,0,&result)

	//决策
	return result
}

func f2(target int,choice []int,temp []int,index int,result *[][]int) {

	//判断终止
	if target < 0 {
		return
	}

	//判断是否我相等
	if target == 0 {

		*result = append(*result,temp)
		return
	}
	//增加缓存
	mp := map[int]int{}

	//决策
	for i := index; i < len(choice) && choice[i] <= target; i++{//满足条件的决策

	  //判断是否出现过
	   if choice[i] == 0 || mp[choice[i]] == 1{

	   		continue
	   }

	   //由于只能出现一次
		tValue := choice[i]

		//赋值为0表示已经取过了
		choice[i] = 0

		//深拷贝结果
		tmp := make([]int,len(temp))
		copy(tmp,temp)

		//记录决策
		tmp = append(tmp,tValue)

		//决策下一步
		f2(target - tValue,choice,tmp,i,result)

		//恢复
		choice[i] = tValue

		//写入历史
		mp[tValue] = 1

	}

}

func main() {
	arr := []int{10,1,2,7,6,1,5}

	fmt.Println(combinationSum2(arr,8))

}
