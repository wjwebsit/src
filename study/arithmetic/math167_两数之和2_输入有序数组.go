package main
/**
	算法描述:
	给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func twoSum(numbers []int, target int) []int {
	//声明返回数组
	var result []int

	//判断极端情况
	if len(numbers) < 2  {
		return result
	}
	if numbers[0] >= target {
		return  result
	}

	//采用双指针
	start,end := 0,len(numbers) - 1
	for start < end {
		//获取和
		sum := numbers[start] + numbers[end]

		//判断目标和
		if sum > target {
				end --
		} else if sum < target {
				start ++
		} else {
				//写入结果
				result = append(result,start + 1)
				result = append(result,end +1)
				break
		}

	}

	//返回结果
	return result
}

func main() {

}
