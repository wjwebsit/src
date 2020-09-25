package main

import "math"

/**
	算法描述
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
 */
func permuteUnique(nums []int) [][]int {
	//定义返回值
	var result [][]int

	//判断
	if len(nums) == 0 {
		return result
	}

	tmp := []int{}

	f4 (nums,tmp,0,&result)


	//返回
	return result

}

func f4 (nums []int, tmp []int,count int ,result *[][]int) {
	//判断结束
	if count == len(nums) {
		*result = append(*result,tmp)
		return
	}
	//加入缓存
	mp := map[int]int{}

	//决策
	for i := 0 ; i < len(nums); i ++ {
		//pan断
		if nums[i] == math.MaxInt64 || mp[nums[i]] == 1{
			continue
		}

		//选择
		tValue := nums[i]
		nums[i] = math.MaxInt64   //设置为无穷 下次就不会在选它

		//深拷贝结果
		tmp1 := make([]int,len(tmp))
		copy(tmp1,tmp)

		//写入结果
		tmp1 = append(tmp1,tValue)

		//下一次
		f4(nums,tmp1,count + 1,result)

		//恢复
		nums[i] = tValue

		//写入缓存
		mp[tValue] = 1
	}
}