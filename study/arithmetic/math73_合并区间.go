package main

import "fmt"

/**
	算法描述:
	给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 */
/**
	动态规划dp
 */
func mergeQujian(intervals [][]int) [][]int {
	//判断条件
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	//区间排序
	for i := 1; i < len(intervals); i++ {
		//获取当前
		interval := intervals[i]

		//判断
		j := i - 1
		for j >= 0  && (intervals[j][0] > interval[0] || (intervals[j][0] == interval[0] && interval[len(interval) - 1] < intervals[j][len(intervals[j]) - 1])) {
			intervals[j + 1] = intervals[j]
			j --
		}

		intervals[j + 1] = interval
	}

	//动态规划
	var dp [][]int

	//填入第一项
	dp = append(dp,intervals[0])

	//开始
	for i := 1; i < len(intervals);i++ {
		dp = mergeSection(dp,intervals[i])
	}

	//返回
	return dp

}

func mergeSection(dp [][]int,arr2 []int)[][]int {
	//判断
	if len(arr2) == 0 {
		return dp
	}

	//获取长度
	length := len(dp)

	//判断
	if length == 0 {
		dp = append(dp,arr2)
		return dp
	}

	//判断
	if dp[length - 1][len(dp[length - 1]) - 1] < arr2[0] {
		dp = append(dp,arr2)
		return dp
	}

	//合并
	var temp []int
	temp = append(temp,dp[length - 1][0])

	//另一端
	if dp[length - 1][len(dp[length - 1]) - 1] > arr2[len(arr2) - 1] {
		temp = append(temp,dp[length - 1][len(dp[length - 1]) - 1])
	} else {
		temp = append(temp,arr2[len(arr2) - 1])
	}

	dp  = dp[:length - 1]
	dp = append(dp,temp)

	//返回
	return dp
}

func main() {
	var intervals [][]int
	intervals = append(intervals,[]int{5,7})
	intervals = append(intervals,[]int{5,5})
	intervals = append(intervals,[]int{1,1})
	intervals = append(intervals,[]int{0,0})
	intervals = append(intervals,[]int{3,3})
	intervals = append(intervals,[]int{4,5})
	intervals = append(intervals,[]int{1,1})
	intervals = append(intervals,[]int{3,4})

	fmt.Println(mergeQujian(intervals))

}



