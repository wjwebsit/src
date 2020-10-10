package main

import "fmt"

/**
	算法描述
	给你一个整数数组 arr ，以及 a、b 、c 三个整数。请你统计其中好三元组的数量。

如果三元组 (arr[i], arr[j], arr[k]) 满足下列全部条件，则认为它是一个 好三元组 。

0 <= i < j < k < arr.length
|arr[i] - arr[j]| <= a
|arr[j] - arr[k]| <= b
|arr[i] - arr[k]| <= c
其中 |x| 表示 x 的绝对值。

返回 好三元组的数量 。

 

示例 1：

输入：arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
输出：4
解释：一共有 4 个好三元组：[(3,0,1), (3,0,1), (3,1,1), (0,1,1)] 。
示例 2：

输入：arr = [1,1,2,2,3], a = 0, b = 0, c = 1
输出：0
解释：不存在满足所有条件的三元组。
 

提示：

3 <= arr.length <= 100
0 <= arr[i] <= 1000
0 <= a, b, c <= 1000

 */
func countGoodTriplets(arr []int, a int, b int, c int) int {
	//验证参数
	if len(arr) < 3 {
		return 0
	}

	//A[i] A[j] A[k] 满足 i < j < k
	hash1 := make([]map[int]bool,len(arr) - 1)
	hash2 := make([]map[int]bool,len(arr))
	hash3 := make([]map[int]bool,len(arr))

	//预处理
	for i := 0; i < len(arr) - 1;i ++ {
		//初始化
		hash1[i] = make(map[int]bool)
		hash2[i] = make(map[int]bool)
		hash3[i] = make(map[int]bool)

		for j := i + 1; j < len(arr); j ++ {
			//获取绝对值
			abs := myAbs(arr[i],arr[j])
			if abs <= a {
				hash1[i][j] = true
			}
			if abs <= b {
				hash2[i][j] = true
			}
			if abs <= c && j - i >= 2 {
				hash3[i][j] = true
			}
		}

	}

	sum := 0
	//统计三元祖
	for i := 0; i < len(arr) - 1;i ++ {
		//获取hash1
		if len(hash1[i]) == 0 {
			continue
		}
		for k,_ := range hash1[i] {
			//获取hash2
			for j, _ := range hash2[k] {
				//判断
				if ok, _ := hash3[i][j]; ok {
					sum++
				}
			}
		}
	}

	//返回
	return sum
}

/**
	返回绝对值 a,b > 0
 */
func myAbs(a,b int) int {
	//大数-小数
	if a > b {
		return  a - b
	}
	return b - a
}

func main() {
	arr := []int{3,0,1,1,9,7}

	fmt.Println(countGoodTriplets(arr,7,2,3))
}