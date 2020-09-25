package main

import "fmt"

/**
	算法描述:
	假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0

 */

func findMin(nums []int) int {
	//利用二分查找
	var min = nums[0]

	//双指针
	s,e := 0,len(nums) - 1
	for s <= e {
		if nums[s] == nums[e] {
			//顺序查找
			for s <= e {
				if nums[s] < min {
					min = nums[s]
				}
				s++
			}
			break
 		}

		//获取mid
		mid := (s + e) / 2

		//判断区间连续
		if nums[s] <= nums[mid] {//左区间连续
			//比较左边最小
			if nums[s] < min {
				min = nums[s]
			}
			//查右边
			s = mid + 1

		} else {//右区间连续
			//比较
			if nums[mid] < min {
				min = nums[mid]
			}
			//查左边
			e = mid - 1
		}

	}

	//返回
	return min

}
func main() {
	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(findMin(nums))
}