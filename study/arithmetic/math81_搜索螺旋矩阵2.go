package main

import "fmt"

/**
	算法描述:
	假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false

 */
func search2(nums []int, target int) bool {
	//判断数组
	if len(nums) == 0 {
		return false
	}

	//二分查找
	s,e := 0,len(nums) - 1
	for s <= e {
		//判断是否相等
		if nums[s] == nums[e] {
			//顺序查找
			for s <= e {
				if nums[s] == target {
					return true
				}
				s ++
			}
			break
		}

		//获取mid
		mid := (s + e) / 2

		//判断
		if nums[mid] == target {
			return true
			break
		}

		//左区间连续
		if nums[s] <= nums[mid] {
			if nums[s] <= target && target < nums[mid] {
				e = mid - 1
			} else {
				s = mid + 1
			}

		} else {
			if nums[mid] < target && target <= nums[e] {
				s = mid + 1
			} else {
				e = mid - 1
			}
		}


	}

	return false

}
func main() {
	nums := []int{2,5,6,0,0,1,2}
	fmt.Println(search2(nums,6))
}