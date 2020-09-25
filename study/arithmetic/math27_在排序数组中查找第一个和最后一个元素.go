package main

import "fmt"

/**
	算法描述：
给定一个按照升序排列的整数数组nums，和一个目标值target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是  O（log n）级别。

如果数组中不存在目标值，返回  [-1, -1]。

示例1：

输入： nums = [ 5,7,7,8,8,10]，target = 8
 输出： [3,4]
示例2：

输入： nums = [ 5,7,7,8,8,10]，target = 6
 输出： [ - 1，-1]
 */

 /**
 	典型的折半查找 时间复杂度 log2(n)
  */
func searchRange(nums []int, target int) []int {
	//判断数组长度<=1 情况
	if len(nums) == 0 ||(len(nums) == 1 && nums[0] != target){
		return []int{-1,-1}
	}

	//折半查找
	begin := 0
	end := len(nums) - 1

	//判断极端情况
	if nums[begin] > target || nums[end] < target {
		return []int {-1,-1}
	}

	//定义起始下标
	firstIndex,lastIndex := -1,-1
	targetIndex := -1  //折半查找到的

	//双指针begin和end 适用于len >=2 故判断
	if len(nums) == 1 && nums[0] == target {
		return []int{0,0}
	}

	//开始寻找值
	for begin <= end {
		//定义下标
		mid := (begin + end) /2

		//缩小搜索域
		if nums[mid] < target {
			begin = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			targetIndex = mid
			break
		}

	}

	//判断targetIndex
	if targetIndex == -1 {
		return []int {-1,-1}
	}

	//寻找最早和最晚
	for i := targetIndex; i >=0 && nums[i] == target; i-- {
		firstIndex = i
	}
	for j := targetIndex;j < len(nums) && nums[j] == target;j ++ {
		lastIndex = j
	}

	//返回值
	return []int{firstIndex,lastIndex}

}

//测试
func main () {
	arr := []int{1,2,3,4}

	fmt.Println(searchRange(arr,1))
}