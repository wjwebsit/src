package main
/**
	算法描述
给定一个包含非负数的数组和一个目标整数 k，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为 2，总和为 k 的倍数，即总和为 n*k，其中 n 也是一个整数。

示例 1:

输入: [23,2,4,6,7], k = 6
输出: True
解释: [2,4] 是一个大小为 2 的子数组，并且和为 6。
示例 2:

输入: [23,2,6,4,7], k = 6
输出: True
解释: [23,2,6,4,7]是大小为 5 的子数组，并且和为 42。
 */
 /**
 	Arr(i,j) = Arr[j] - arr[i - 1]
 	for i < j  Arr[j] - arr[i] = n * k n >= 1  => Arr[j] - n * k = arr[i]

  */
func checkSubarraySum(nums []int, k int) bool {
	//判断数组
	if len(nums) == 0 {
		return false
	}
	if k < 0 {
		k = -1 * k
	}

	//构造前缀和
	sumArr := make([]int,len(nums) + 1)
	sumArr[0] = 0

	//构造前缀数组
	for i := 1;i <= len(nums); i ++ {
		sumArr[i] = sumArr[i - 1] + nums[i - 1]
	}

	//开始寻找
	for i := 2; i < len(sumArr); i ++ {
		for j := 1; j <= i - 1 ; j ++ {
			//获取和
			sum := sumArr[i] - sumArr[j - 1]

			//判断
			if k == 0 && sum == 0 {
				return true
			}

			if sum % k == 0 {
				return true
			}

		}
	}

	return false


}
func main() {

}