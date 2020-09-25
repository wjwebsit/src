package main

import (
	"fmt"
	"math"
)

/**
	算法描述:
	给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

 

进阶：

你能在线性时间复杂度内解决此题吗？

 

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 

提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
 */
func maxSlidingWindow(nums []int, k int) []int {
	//判断
	if len(nums) == 0 {
		return []int{}
	}


	//定义mp--记录当前滑动窗口最大值下标
	mp := make(map[int]int)

	//获取第一个滑动窗口最大值下标
	max := math.MinInt64
	for i := 0; i < k ; i ++ {
		if nums[i] > max {
			mp[nums[i]] = i
			max = nums[i]
		}
	}
	//声明返回值
	var result []int
	result = append(result,max)
	for i := k; i < len(nums); i ++ {
		//判断要弹出的元素是否越过最大值索引
		if i - k + 1 > mp[max] {
			//寻找最大值
			max = math.MinInt64
			for j := i - k + 1; j <= i; j ++ {
				if nums[j] > max {
					max = nums[j]
					mp[max] = j
				}
			}

			//写入
			result = append(result, max)

		} else {
			//判断当前是否大于等于
			if nums[i] >= max {
				//更新
				max = nums[i]
				mp[max] = i
			}

			//写入
			result = append(result, max)
		}
	}

	//返回结果
	return  result
}
func main() {
	nums := []int {7,2,4}
	fmt.Println(maxSlidingWindow(nums,2))
}
