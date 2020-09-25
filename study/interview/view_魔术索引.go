package main
/**
	算法描述
	魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。若有多个魔术索引，返回索引值最小的一个。

示例1:

 输入：nums = [0, 2, 3, 4, 5]
 输出：0
 说明: 0下标的元素为0
示例2:

 输入：nums = [1, 1, 1]
 输出：1
提示:

nums长度在[1, 1000000]之间
 */
func findMagicIndex(nums []int) int {
	//声明返回
	min := -1

	//折半查找
	s,e := 0,len(nums) - 1

	for s <= e {
		mid := s + (e - s) / 2

		//判断
		if nums[mid] == mid {
			min = mid

			//向前找更小的
			e = mid - 1

		} else if nums[mid] > mid {//值比索引大--后边的均是
			e = mid - 1

		} else {//值索引小
			s = mid + 1
		}
	}
	return min

}
