package main

import (
	"fmt"
	"strconv"
)

/**
	算法描述
	给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。

示例 1:

输入: [0,1,2,4,5,7]
输出: ["0->2","4->5","7"]
解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。
示例 2:

输入: [0,2,3,4,6,8,9]
输出: ["0","2->4","6","8->9"]
解释: 2,3,4 可组成一个连续的区间; 8,9 可组成一个连续的区间。
 */
func summaryRanges(nums []int) []string {
	//声明返回
	var result []string

	//判断
	if len(nums) == 0 {
		return  result
	}

	var i = 0
	for i < len(nums) {
		tmp := strconv.Itoa(nums[i])
		j := i + 1
		for j < len(nums) && nums[j] == nums[j - 1] + 1 {//区间连续
			j ++
		}
		//写入结果
		if j == i + 1 {
			result = append(result, tmp)
		} else {
			result = append(result, tmp + "->" + strconv.Itoa(nums[j - 1]))
		}

		//赋值
		i = j
	}
	//返回
	return result

}

func main() {
	nums := []int{8,9}
	fmt.Println(summaryRanges(nums))
}