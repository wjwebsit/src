package main

import "fmt"

/**
	算法描述
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
 */
  /**
 	回溯算法
  */
func subsets(nums []int) [][]int {
	//声明返回
	var result [][]int

	//判断
	if len(nums) == 0 {
		return result
	}

	for i := 1; i <= len(nums); i++ {
		tmp := []int{}
		f10(i,nums,0,0,tmp,&result)
	}

	//加入[]
	result = append(result,[]int{})

	//返回
	return result



}

/**
	回溯
 */
func f10(k int,choice []int,height int,start int,tmp []int,result *[][]int) {
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
		f10(k,choice,height + 1,start + 1,tmp1,result)

		start ++
	}

}
func main() {
	arr := []int{1,2,3}
	fmt.Println(subsets(arr))
}


