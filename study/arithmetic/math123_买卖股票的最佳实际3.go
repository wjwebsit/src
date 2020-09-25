package main

import "fmt"

/**
	算法描述:
	给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [3,3,5,0,0,3,1,4]
输出: 6
解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。  
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。  
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:

输入: [7,6,4,3,1]
输出: 0
解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
 */

func maxProfit3(prices []int) int {
	 //判断数组
	if len(prices) == 0 {
		return 0
	}

	//利用单调栈
	var stack []int
	var profit []int  //降序排列
	prices = append(prices,-1) //单调栈保证最后出完
	profit = append(profit,0) //添加哨兵减少切片长度为0的判断
	stack = append(stack,prices[0])
	for i := 1; i < len(prices);i ++ {
		//判断
		if prices[i] >= stack[len(stack) - 1] {
			stack = append(stack,prices[i])
			continue
		}

		//判断
		if len(stack) > 1 {
			//计算利润并排序
			num := stack[len(stack) - 1] - stack[0]

			//折半查找插入位置
			low,end := 0,len(profit) - 1
			for low <= end {
				mid := (low + end) / 2
				if profit[mid] <= num {
					end = mid - 1
				} else {
					low = mid + 1
				}
			}

			//此时low即为要插入位置
			if low == len(profit) {
				profit = append(profit,num)
			} else {
				profit = append(profit,0)

				//右移
				for k := len(profit) -2;k >= low;k -- {
					profit[k + 1] = profit[k]
				}
				profit[low] = num
			}
			fmt.Println(profit)
		}
		//栈清空
		stack = []int{}
		stack = append(stack,prices[i])

	}

	//返回
	if len(profit) == 1 {
		return profit[0]
	} else {
		return profit[0] + profit[1]
	}
}
func main() {
	nums := []int{1,2,4,2,5,7,2,4,9,0}
	fmt.Println(maxProfit3(nums))
}
