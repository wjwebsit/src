package main

import "fmt"

/**
	算法描述：
	给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:
输入: nums = [1], k = 1
输出: [1]
 */
/**
 */
func topKFrequent(nums []int, k int) []int {
	//声明散列
	hash := make(map[int]int)

	//遍历数组--统计次数
	for i := 0; i < len(nums); i ++ {
		hash[nums[i]] += 1
	}

	//声明返回
	var result []int

	//遍历hash--构造反hash
	yHash := make(map[int][]int)
	//对其出现次数进行计数排序O（n）
	var count = make([]int,len(nums) + 1)//出现次数最大值< 数组长度
	for key,v := range hash {
		//出现次数数组
		count[v] += 1
		yHash[v] = append(yHash[v],key)
	}

	for i := len(count) - 1; i >= 0 && len(result) < k ;i -- {
		//判断
		if count[i] == 0 {
			continue
		}
		result = append(result,yHash[i]...)
	}


	//返回
	return result
}
func main() {
	fmt.Println(topKFrequent([]int{1,2},2))

}