package main

import "fmt"

/**
	算法描述：
	给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	//声明辅助数组---注意为深拷贝
	tmpNums := make([]int,m)
	copy(tmpNums,nums1[:m])

	//采用双指针 i指向tmpNums j指向nums2
	i,j,k := 0,0,0

	//遍历邻接条件
	for i < m || j < n {
		//判断
		for i == m && j < n {
			nums1[k] = nums2[j]
			j++
			k++
		}

		//判断
		for j == n && i < m {
			nums1[k] = tmpNums[i]
			i++
			k++
		}

		//判断
		for i < m && j < n && tmpNums[i] <= nums2[j] {
			nums1[k] = tmpNums[i]
			i++
			k++
		}

		//判断
		for i < m && j < n  && tmpNums[i] > nums2[j] {
			nums1[k] = nums2[j]
			j++
			k++
		}

	}

	fmt.Println(nums1)

}

func main() {
	arr1 := []int{1,2,3,0,0,0}
	arr2 := []int{2,5,6}

	merge(arr1,3,arr2,2)
}
