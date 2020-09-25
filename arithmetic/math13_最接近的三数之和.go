package main

import (
	"fmt"
	"math"
	"sort"
)

/**
	算法描述
给定一个包括  n个整数的数组  nums 和一个目标值  target。找出  nums 中的三个整数，使得它们的和与  target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组nums = [-1,2,1，-4]，和target = 1。

与目标最接近的三个数的和为2.（-1 + 2 + 1 = 2）。
 */

 func main() {
 	nums := []int{-55,-24,-18,-11,-7,-3,4,5,6,9,11,23,33}

	fmt.Print(threeSumClosest(nums,0))
 }
/**
	双指针法寻找最接近剩余target
 */
func threeSumClosest(nums []int, target int) int {
	//判断
	if len(nums) < 3 {
		return -1
	}

	//排序数组
	sort.Ints(nums)

	//最接近的和
	var sum int

	//最接近
	for i := 0; i <= len(nums) - 3 ; i ++ {
		j := i + 1
		k := len(nums) - 1

		//判断是否相等
		if nums[i] + nums[j] + nums[k] == target {
			sum = target
			break
		}
		//上一个最接近的比target还要大 那后面的更大 更不接近
		/*if target > 0 && i != 0 && sum > target {
			continue
		}*/

		//判断sum是否相等
		if i != 0 && sum == target {

			break;
		}
		//接下来最接近
		lastTarget := target - nums[i]

		//由于接近为比其大和比其小故选取临界--用数组记录
		limit := []int{}

		for j < k {
			sumIk := nums[j] + nums[k]
			//判断是否相等
			if sumIk == lastTarget {
				sum = target
				break
			} else if sumIk < lastTarget { //记录最次比target小的时候
				j ++
				//判断是否是最后
				if (j < k && nums[j] + nums[k] > lastTarget) || j >= k{
					limit = append(limit,sumIk)
				}

			} else { //记录最后一次比target大的时候
					k --
				if (j < k && nums[j] + nums[k] < lastTarget) || k <= j {
					limit = append(limit,sumIk)
				}
			}
		}

		//判断临界
		tempSum := 0


		//结束；
		if len(limit) == 0 {
			continue
		}


		if len(limit) == 1 {
			tempSum = nums[i] + limit[0]

		} else if len(limit) >= 2 {
			if math.Abs(float64(lastTarget - limit[0])) < math.Abs(float64(lastTarget - limit[1])) {
				tempSum = limit[0] + nums[i]
			} else {
				tempSum = limit[1] + nums[i]
			}
		}

		//和sum比较
		if i==0 {
			sum = tempSum
		} else if math.Abs(float64(target - sum)) > math.Abs(float64(target - tempSum)) {
			sum = tempSum
		}



	}

	//返回
	return sum


}