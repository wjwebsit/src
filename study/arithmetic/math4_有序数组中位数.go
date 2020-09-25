package main

import "fmt"

/**
给定两个大小为m和n的有序数组  nums1和  nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为O（log（m + n））
示例1：

nums1 = [1,3]
nums2 = [2]

则中位数是2.0
示例2：

nums1 = [1,2]
nums2 = [3,4]

则中位数是（2 + 3）/ 2 = 2.5
 */
func main() {
	nums1 := []int{1,2}
	nums2 := []int{3,4}

	fmt.Print(medianArray(nums1,nums2))
}

func medianArray(arr1 []int,arr2 []int) float64 {
	//获取切片长度
	arrL1 := len(arr1)
	arrL2 := len(arr2)

	sumL := arrL1 + arrL2

	//sumL为奇数或偶数中位数的数量不同
	var medianStart, medianEnd int //奇数 mediaStart = medianEnd

	//根据数组长度来确认中位数位置
	if sumL%2 == 1 {
		medianStart = int(sumL / 2)
		medianEnd = medianStart
	} else {
		medianStart = int(sumL/2) - 1
		medianEnd = medianStart + 1
	}

	//定义数组的索引
	index1 := 0
	index2 := 0
	count := 0 //整个串的位置

	//定义位数集
	var result []int

	//执行
	for count <= medianEnd {
		for index1 == len(arr1) && index2 < len(arr2) && count <= medianEnd {
			if count == medianStart || count == medianEnd {
				result = append(result, arr2[index2])
			}

			index2++
			count++
		}
		for index2 == len(arr2) && index1 < len(arr1) && count <= medianEnd {
			if count == medianStart || count == medianEnd {
				result = append(result, arr1[index1])
			}

			index1++
			count++
		}

		for index1 < len(arr1) && index2 < len(arr2) && count <= medianEnd && arr1[index1] >= arr2[index2] {
			if count == medianStart || count == medianEnd {
				result = append(result, arr2[index2])
			}

			index2 ++
			count ++
		}

		for index2 < len(arr2) && index1 < len(arr1) && count <= medianEnd && arr2[index2] >= arr1[index1] {
			if count == medianStart || count == medianEnd {
				result = append(result, arr1[index1])
			}
			index1 ++
			count ++
		}

	}

	//返回值
	if sumL %2 == 1 {
		return float64(result[0])
	} else {

		return (float64(result[0]) + float64(result[1]))/2
	}

}
