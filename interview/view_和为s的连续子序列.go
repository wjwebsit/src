package main
/**
	算法描述
	输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

 

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
 

限制：

1 <= target <= 10^5

 */
func findContinuousSequence(target int) [][]int {
	//声明返回值
	var result [][]int

	//判断
	if target <= 0 {
		return result
	}

	//声明hash
	hash := make(map[int]int)
	hash[0] = 0

	//前缀和
	sum := 0
	for i := 1; i <= target / 2 + 1; i ++ {
		sum += i

		//判断
		if _,ok := hash[sum - target]; ok && i - hash[sum - target] >= 2 {//存在且至少包含2个数
			//构造
			var tmp []int
			for j := hash[sum - target]; j <= i ; j ++ {
				tmp = append(tmp,j)

			}

			//写入结果
			result = append(result,tmp)

		}

		//记录
		hash[sum] = i

	}

	//返回
	return result




}