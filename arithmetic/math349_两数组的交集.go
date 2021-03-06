package main
/**
	算法描述
	给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
说明:

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func intersection(nums1 []int, nums2 []int) []int {
	//声明结果
	var result []int

	//判断
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	//声明mp
	mp1,mp2 := make(map[int]int),make(map[int]int)

	//构造数组1的map
	for _,v := range nums1 {
		mp1[v] = 1
	}

	//数组2
	for _,v := range nums2 {
		//相交且不重复
		if mp1[v] == 1 && mp2[v] != 1 {
			//写入
			result = append(result,v)
		}

		//写入
		mp2[v] = 1

	}

	//返回
	return result

}