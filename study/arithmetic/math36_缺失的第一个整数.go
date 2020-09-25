package main

import (
	"fmt"
	"math"
)

/**
	算法描述
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
 */
 /**
 	大神思路 ---类似计数排序
 	遍历一次数组把大于等于1的和小于数组大小的值放到原数组对应位置，然后再遍历一次数组查当前下标是否和值对应，如果不对应那这个下标就是答案，否则遍历完都没出现那么答案就是数组长度加1。
  */

 func  firstMissingPositive2(num []int) int {
 	//定义数组
 	var numArr [len(num)] int

 	for _,v := range num {
 		if v >= 1 {
 			numArr[v - 1] = v
		}

	}

 	i := 0

 	for i := 0; i < len(num);i ++ {
 		if numArr[i] != i + 1 {
 			break
		}
	}

 	return i + 1

 }




/**
	尾插数组来实现---不好
 */
func firstMissingPositive(nums []int) int {

	//判断
	if len(nums) == 0 {
		return 1
	}

	//计算长度
	length := len(nums)

	//定义min
	rtMin := 1

	//最小标记为+无穷
	min := math.MaxInt64
	lastMin := min

	//最大
	max := 0

	fg := true

	//循环
	for i := 0 ; i < len(nums);i++{//依次计算与min的差
		//追加末尾--为了界限
	     nums = append(nums,nums[i])


		//将差插入末尾
		if i < length { //寻找最小的2个元素

			if nums[i] < min  && nums[i] > 0{
				min = nums[i] //最小
			}
			if nums[i] > max  && nums[i] > 0{
				max = nums[i] //最大
			}
			//判断是否结束
			if i  == length - 1 {

				//判断
				if min > 1 { //最小不为1 直接返回
					rtMin = 1
					break
				}

				if max == 1 {
					rtMin = 2
					break;
				}


				//寻找比2 大的第一个数
				min = 2
				lastMin = 1

			}


		} else { //表示记录过一次

		    //判断是否比target 大的第一个数
			if nums[i] > lastMin && nums[i] <= min  && (i + 1) % length >= 0{
				fg = false
				min = nums[i]
			}

			//临界
			if (i + 1) % length == 0 {

				if min - lastMin > 1 {
					rtMin = lastMin + 1
					break

				} else if fg ||(min - lastMin == 1 && min == max){
					if fg {
						min --
					}

					rtMin = min + 1
					break
				} else {

					lastMin = min
					min ++
					fg = true
				}

			}
		}
	}

	//返回
	return rtMin
}

func main() {
	arr := []int{1,2,0}
	fmt.Println(firstMissingPositive(arr));

}