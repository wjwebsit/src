package main

import "fmt"

/**
	题目描述：分支限界法
	给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 */
 /**
 	类似于算法导论里的不相交集合的数据结构来求解图的最大连通量
  */
func longestConsecutive(nums []int) int {
	//定义 hash表 key为不相交结合的左右两端 值相同为集合数组和插值(0下标)
	//例如 101 =>[1,101,102] 102=>[1,101,102]
	hash := map[int][]int{}

	//定义返回值
	max := 0

	//遍历数组
	for _,v := range nums {
		//初始化新结合
		var newSet []int

		//当前为不相交集合
		if len(hash[v - 1]) == 0 && len(hash[v + 1]) == 0 && len(hash[v]) == 0{
			//插值为0
			newSet = append(newSet,0)

			//集合值
			newSet = append(newSet,v)

			//写入
			hash[v] = newSet

			//比较
			if max < 1 {
				max = 1
			}
			continue
		}

		//判断v-1是否存在且以v-1为界
		if len(hash[v - 1]) > 0 {
			//获取区间
			tmpSet := hash[v - 1]

			//以v-1为界
			if tmpSet[len(tmpSet) -  1] != v - 1 {
				continue
			}

			//获取区间的上界
			start := v - 1 - tmpSet[0]

			//将区间更新
			tmpSet = append(tmpSet,v)
			tmpSet[0] ++

			//起点更新
			hash[start] = tmpSet

			//hash[v -1 ] 删除  并生成 hash[v]
			if (v - 1) != start {//即[2]，[3] 合并[2,3]无需删除
				delete(hash, v - 1)
			}
			hash[v] = tmpSet

			//更新max
			if max < tmpSet[0] + 1 {
				max = tmpSet[0] + 1
			}
		}
		//判断右区间是否存在且以v+1开始
		if len(hash[v + 1]) > 0 {
			//获取区间
			tmpSet := hash[v + 1]

			//判断
			if tmpSet[1] != v + 1 {
				continue
			}

			//获取右区间的闭区间
			end := v + 1 + tmpSet[0]

			//判断hash[v] 是否存在
			if len(hash[v]) > 0  {//v - 1 和 v + 1 都有情况
				//获取hash [V]区间 start和区间 并去除差值
				start := v - hash[v][0]
				tmpSet1 := hash[v][1:]

				//tmpSet去除差值
				tmpSet = tmpSet[1:]

				//加算新差值
				diff := end - start

				//新的集合
				newSet = append([]int{diff},tmpSet1...)
				newSet = append(newSet,tmpSet...)

				//更新start,end 删除 v,v+1
				delete(hash,v)
				delete(hash,v + 1)
				hash[start] = newSet
				hash[end] = newSet

				//比较max
				if max < diff + 1 {
					max = diff + 1
				}

			} else {
				//v填入tmpSet
				tmpSet = tmpSet[1:]
				tmpSet = append([]int{v},tmpSet...)

				//差值更新 end - v
				tmpSet = append([]int{end - v},tmpSet...)

				//删除v+1
				delete(hash,v + 1)

				//更新end
				hash[end] = tmpSet
				hash[v] = tmpSet

				//比较max
				if max < tmpSet[0] + 1 {
					max = tmpSet[0] + 1
				}
			}

		}

	}

	//返回
	return max
}
func main() {
	nums := []int{3,2,4,4,5,7,8,6,9,11,10}
	fmt.Println(longestConsecutive(nums))
}