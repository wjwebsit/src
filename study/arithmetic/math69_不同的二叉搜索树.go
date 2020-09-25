package main

import "fmt"

/**
	算法描述:
	给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
 */
func numTrees(n int) int {
	//定义sum
	sum := 0

	//判断
	if n <= 0 {
		return sum
	}

	//初始化切片
	var nums []int
	for i := 1; i <= n; i ++ {
		nums = append(nums,i)
	}

	//开始寻找

	for i := 1;i <= n; i ++ {
		nums1 := make([]int,len(nums))
		copy(nums1,nums)
		nums1 = append(nums1[:i - 1],nums1[i:]...)
		sum += f17(nums1,n,1,i,i)
	}

	//返回
	return sum


}
/**
回溯
 */
func f17(nums []int,n int,height int,parent int,root int)int {
	//判断界
	if height == n {
		return 1
	}

	//开始寻找当前可以作为左孩子的 < parent,和右孩子的 >parent
	var left, right []int
	for i,v := range nums {
		if v < parent  && v < root{
			left = append(left,i)
			continue
		}
		if v > parent && v > root {
			right = append(right, i)
		}

	}

	//定义sum
	sum := 0

	//判断
	if len(left) > 0 {
		for _,v := range left {
			//深拷贝
			nums1 := make([]int,len(nums))
			copy(nums1,nums)

			//删除当前
			temp := nums1[v]
			nums1 = append(nums1[0:v],nums1[v + 1:]...)

			sum += f17(nums1,n,height + 1,temp,root)

		}

	}

	if len(right) > 0 {
		for _,v := range right {
			//深拷贝
			nums2 := make([]int,len(nums))
			copy(nums2,nums)

			//删除当前
			temp := nums2[v]
			nums2 = append(nums2[0:v],nums2[v + 1:]...)


			sum += f17(nums2,n,height + 1,temp,root)

		}

	}

	//返回
	return sum
}

func main() {
	fmt.Println(numTrees(3))
}