package main
/**
	算法描述
	给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。

注意:
数组长度 n 满足以下条件:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
示例:

输入:
nums = [7,2,5,10,8]
m = 2

输出:
18

解释:
一共有四种方法将nums分割为2个子数组。
其中最好的方式是将其分为[7,2,5] 和 [10,8]，
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
 */

func splitArray(nums []int, m int) int {
	//获取数组的最大值、以及和
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i ++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}

	//采用二分查找
	s,e := max,sum

	for s <= e {
		mid := s + (e - s) / 2

		//判断是否满足条件
		if guess(nums,mid,m) {
			//寻找更大
			s = mid + 1

		} else {//和太大了
			e = mid
		}

	}

	//返回
	return s
}

func guess(nums []int,target,m int) bool {
	//将nums 分割m 个 且每一个均不小于target
	sum := 0
	count := 0
	for i := 0 ; i < len(nums); i ++ {
		if sum + nums[i] >= target {
			count ++
			sum = 0
		} else {
			sum += nums[i]
		}
	}

	//判断
	return count >= m

}