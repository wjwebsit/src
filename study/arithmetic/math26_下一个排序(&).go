package main

import (
	"fmt"
	"sort"
)

/**
	算法描述:
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

*/

func nextPermutation(nums []int) []int  {
	//判断长度
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	//定义辅助指针
	last := len(nums) - 1
	pre := last - 1

	//倒序遍历----前部保持升序
	for pre >= 0 {
		//last 为前半部分的升序数组最后一个
		if nums[pre] >= nums[last] {
			pre --
			last--

		} else { //表示位遍历完存在---找到合适位置
			index := last


			//寻找最开始大于的
			for index <= len(nums) - 1 && nums[index] > nums[pre] {
				index ++
			}

			nums[index - 1],nums[pre] = nums[pre],nums[index - 1]

			//转换
			nums = reserveIntArr(nums,pre)

			break

		}

	}

	//判断是否为最大
	if pre == -1 {
		sort.Ints(nums)
	}

	//返回
	return nums


}

/**
反转切片
 */
func reserveIntArr( arr []int,startIndex int) []int {
	//定义返回
	var rtArr []int

	//反转数组
	for end := len(arr) - 1 ; end > startIndex; end -- {
		rtArr = append(rtArr,arr[end])
	}

	for i := 0; i< len(rtArr); i++ {
		arr[startIndex + 1 + i] = rtArr[i]
	}

	//返回
	return arr

}

//测试
func main () {
	arr := []int{1,3,2}  // 2,1,3
	fmt.Println(nextPermutation(arr))
}
