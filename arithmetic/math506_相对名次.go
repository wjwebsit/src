package main

import (
	"fmt"
	"strconv"
)

/**
	算法描述
	给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold Medal", "Silver Medal", "Bronze Medal"）。

(注：分数越高的选手，排名越靠前。)

示例 1:

输入: [5, 4, 3, 2, 1]
输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
解释: 前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal" and "Bronze Medal").
余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。
提示:

N 是一个正整数并且不会超过 10000。
所有运动员的成绩都不相同。

 */
func findRelativeRanks(nums []int) []string {
	//构造索引映射
	hash := make(map[int]int)
	for i := 0; i< len(nums);i++ {
		hash[nums[i]] = i
	}

	//构造大顶堆
	for i := (len(nums) - 1) / 2; i >= 0 ;i--{
		nums = heapAdjust(nums,i,len(nums) - 1)
	}


	nHash := map[int]string{
		1:"Gold Medal",
		2:"Silver Medal",
		3:"Bronze Medal",
	}
	//返回
	var buffer = make([]string,len(nums))
	index := 1 //排名
	for i := 0; i < len(nums);i ++ {
		//取出排名
		var por = strconv.Itoa(index)
		if _,ok := nHash[index];ok {
			por = nHash[index]
		}
		buffer[hash[nums[0]]] = por
		index ++

		//调整
		nums[0],nums[len(nums) - 1 - i] = nums[len(nums) - 1 - i],nums[0]
		nums = heapAdjust(nums,0,len(nums) - 2 - i)


	}

	return buffer

}

/**
	调节大顶堆
 */
func heapAdjust(A []int,s,e int) []int  {
	//计算当前的key
	key := A[s]

	for i := 2*s + 1; i <= e; i = 2*s + 1 {
		//找一个在较大的
		if i + 1 <= e && A[i + 1] > A[i] {
			i ++
		}

		//比较
		if A[i] <= key {
			break
		}
		//记录
		A[s] = A[i]
		s = i
	}
	A[s] = key
	return A
}
func main() {
	nums := []int{10,3,8,9,4}
	fmt.Println(findRelativeRanks(nums))
}