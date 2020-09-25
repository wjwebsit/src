package main
/**
	算法描述:
	给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1:

输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
示例 2:

输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 */

/**
	这个貌似比合并区间简单啊
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	//插入
	intervals = append(intervals,newInterval)
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
		dp = section(dp,intervals[i])
	}

	//返回
	return dp

}
func section(dp [][]int,arr2 []int)[][]int {
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