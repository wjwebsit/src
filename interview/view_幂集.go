package main

import "fmt"

/**
	算法描述
	幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。

说明：解集不能包含重复的子集。

示例:

 输入： nums = [1,2,3]
 输出：
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
 */
func subsets(nums []int) [][]int {
	//声明返回
	var result [][]int

	//添加空
	result = append(result,[]int{})

	//判断
	if len(nums) == 0 {
		return result
	}

	for i := 1; i <= len(nums); i ++ {
		DfsSubset(nums,0,0,i,[]int{},&result)
	}

	//返回
	return result


}

/**
	回溯
 */
 func DfsSubset(nums []int,begin,cur,sum int,tmp []int,result *[][]int) {
 	//判断
 	if sum == cur {
 		//写入结果集
 		tmp2 := make([]int,len(tmp))
 		copy(tmp2,tmp)

		*result = append(*result,tmp2)
		return
	}

 	//不够
 	if begin + sum - cur > len(nums) {
 		return
	}

 	//开始
 	for i := begin ; i < len(nums); i ++ {
 		//添加
 		tmp = append(tmp,nums[i])

 		//dfs
 		DfsSubset(nums,i + 1,cur + 1,sum,tmp,result)

 		//回溯
 		tmp = tmp[:len(tmp) - 1]
	}

 }
func main() {
	nums := []int{1,2,3}
	fmt.Println(subsets(nums))
}