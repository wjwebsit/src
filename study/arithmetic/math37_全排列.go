package main

import (
	"fmt"
	"math"
)

/**
	算法描述:
	给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
 */

func permute(nums []int) [][]int {
	//定义返回值
	var result [][]int

	//判断
	if len(nums) == 0 {
		return result
	}

	tmp := []int{}

	f3 (nums,tmp,0,&result)


	//返回
	return result


}

func f3 (nums []int, tmp []int,count int ,result *[][]int) {
	//判断结束
	if count == len(nums) {
		*result = append(*result,tmp)
		return
	}
	//决策
	for i := 0 ; i < len(nums); i ++ {
		//pan断
		if nums[i] == math.MaxInt64 {
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
		f3(nums,tmp1,count + 1,result)

		//恢复
		nums[i] = tValue
	}
}

func main() {
	arr := []int{1,2,3}

	fmt.Println(permute(arr))
}
