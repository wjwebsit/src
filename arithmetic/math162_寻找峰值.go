package main

import (
	"fmt"
	"math"
)

/**
	算法描述:
	峰值元素是指其值大于左右相邻值的元素。

给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。

数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞。

示例 1:

输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。
示例 2:

输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。

要求：时间复杂度(O(logN))
 */

func findPeakElement(nums []int) int {
	//判断
	if len(nums) == 0 {//没有峰值
		return -1
	}

	//补齐数组 ---头尾补全
	nums = append(nums,math.MinInt64)
	nums = append([]int{math.MinInt64},nums...)

	//left,&& right
	left := 2
	right := len(nums) - 3


	//采用滑动窗口每次三个---从左往右，从右往左2个窗口同时进行--do-while
	for true {
		//判断中间元素

		if nums[left - 1] > nums[left]{//中间da于右边
			if nums[left - 1] > nums[left - 2] {
				return left - 1 - 1 //中间的由于补了负无穷
			} else {//向右滑动2个窗口
				left += 2
			}

		} else {
			//向右滑动1
			left += 1
		}
		if nums[right + 1] > nums[right] {
			if nums[right + 1] > nums[right + 2] {
				return right
			} else {//向左滑动2个窗口
				right -= 2
			}

		} else {
			//向左滑动1
			right -= 1
		}

		if (left - 2) > right {//窗口重叠
			break
		}
	}
	return  -1
}
func main()  {
	nums:=[]int{1,2,3,1}
	fmt.Println(findPeakElement(nums))
}