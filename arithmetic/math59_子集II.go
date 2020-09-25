package main

import (
	"fmt"
	"sort"
)

/**
	算法描述:
	给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
 */

func subsetsWithDup(nums []int) [][]int {
	//声明返回
	var result [][]int

	//判断
	if len(nums) == 0 {
		return result
	}
	//排序
	sort.Ints(nums)

	for i := 1; i <= len(nums); i++ {
		tmp := []int{}
		f12(i,nums,0,0,tmp,&result)
	}

	//加入[]
	result = append(result,[]int{})

	//返回
	return result
}
/**
	回溯
 */
func f12(k int,choice []int,height int,start int,tmp []int,result *[][]int) {
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

	//添加缓存
	mp := map[int]int{}

	//遍历
	for start < len(choice) && height + 1 <= k{
		if  mp[choice[start]] == 1 {
			start++
			continue
		}
		//深拷贝
		tmp1 := make([]int,len(tmp))
		copy(tmp1,tmp)

		//写入当前
		tmp1 = append(tmp1, choice[start])

		//写入缓存
		mp[choice[start]] = 1

		//继续决策
		f12(k,choice,height + 1,start + 1,tmp1,result)

		//向下
		start ++
	}

}

func main() {
	arr := []int{4,4,4,1,4}
	fmt.Println(subsetsWithDup(arr))
}
