package main
/**
	算法描述
	给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。

初始化 A 和 B 的元素数量分别为 m 和 n。

示例:

输入:
A = [1,2,3,0,0,0], m = 3
B = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
说明:

A.length == n + m
 */
func merge(A []int, m int, B []int, n int) {
	//构造返回结果
	result := make([]int,m + n)

	//声明下标
	i,j, k := 0,0,0

	for i < m && j < n {
		//比较
		for i < m && A[i] <= B[j] {
			//写入
			result[k] = A[i]

			//next
			k ++
			i ++
		}
		for j < n && A[j] < B[i] {
			//同上
			result[k] = B[j]

			j ++
			k ++
		}

	}

	//添加最后
	if i < m {
		for i < m {
			result[k] = A[i]
			i++
		}

	}

	if j < n {
		for j < n {
			result[k] = B[j]
			j ++
		}
	}
	copy(A,result)



}