package main

import (
	"fmt"
	"sort"
)

/**
	算法描述：
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */

 func main() {
	nums := []int {1,-2,-5,-4,-3,3,3,5}
	fmt.Print(fourSum2(nums,-11))
 }
/**
	四个数之和
 */
func fourSum(nums []int, target int) [][]int {
	//构造返回数组
	result := [][]int{}

	//判断长度
	if len(nums) < 4 {
		return result
	}
	//排序
	sort.Ints(nums)

	//遍历
	for i := 0; i <= len(nums) - 4 ; i++ {//最外一层

		//
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		//相当于三数之和
		for j := i + 1 ; j <= len(nums) - 3; j ++ {

			//同理判断
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}

			//采用双指针法
			begin := j + 1
			end := len(nums) - 1


			for begin < end {

				//判断
				if nums[i] + nums[j] + nums[begin] + nums[end] == target {
					//存放结果
					temp := []int{}
					temp = append(temp,nums[i],nums[j],nums[begin],nums[end])
					//放入结果集
					result = append(result,temp)

					begin ++
					for begin < end && nums[begin] == nums[begin - 1] {
						begin ++
					}

				} else if nums[i] + nums[j] + nums[begin] + nums[end] > target {
					end --
				} else {
						begin ++
				}
			}

		}


	}

	//返回
	return result

}
func fourSum2(nums []int, target int) [][]int{
	//构造返回数组
	result := [][]int{}

	//判断长度
	if len(nums) < 4 {
		return result
	}
	//排序
	sort.Ints(nums)

	fourSum2HS(nums,4,0,target,0,[]int{},&result)

	//返回
	return result


}

func fourSum2HS(nums []int,k int,start int,target int,count int,tmpRes []int,result *[][]int) {
	//判断
	if count == k - 2 {
		end := len(nums) - 1

		for start < end {
			sum := nums[start] + nums[end]

			if sum < target {
				start ++
			} else if sum > target {
					end --
			} else {
				//写入
				var tmp = make([]int,len(tmpRes))
				copy(tmp,tmpRes)
				tmp = append(tmp,nums[start])
				tmp = append(tmp,nums[end])
				(*result) =append((*result),tmp)
				break

			}
		}

		return
	}

	//遍历
	mp := make(map[int]int)
	for i := start; i <= (len(nums) - k + count); i ++ {

		//判断当前如果比target大
		if (target > 0 && nums[i] > target) || (target < 0 && nums[i] > 0){
			break  //后面的都大
		}

		//判断是否重复
		if mp[nums[i]] == 1 {
			continue
		} else {
			//写入
			mp[nums[i]] = 1
		}


		tmpRes = append(tmpRes,nums[i])

		//回溯
		fourSum2HS(nums,k,i + 1,target - nums[i],count + 1,tmpRes,result)

		//回溯
		tmpRes = tmpRes[:len(tmpRes) - 1]

	}

}
