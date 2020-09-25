package main

import "fmt"
/**
	dynamic programming
	动态规划有两种等价的实现方法。

（1）带备忘的自顶向下法

此方法仍按自然的递归形式编写过程，但过程会保存每个子问题的解（通常保存在一个数组或散列表中）。当需要一个子问题的解时，过程首先检查是否已经保存过此解。如果是，则直接返回保存的值，从而节省了计算时间；否则，按通常方式计算这个子问题。

（2）自底向上法

这种方法一般需要恰当定义子问题“规模”的概念，使得任何子问题的求解都只依赖于“更小的”子问题的求解。因此，我们可以将子问题按照规模顺序，由小至大顺序进行求解。当求解某个子问题时，它所依赖的那些更小的子问题都已求解完毕，结果已经保存。每个子问题只需求解一次，当我们求解它时，它的所有前提子问题都已求解完成。

说明：两种方法得到的算法具有相同的渐进运行时间，仅有的差异是在某些特殊情况下，自顶向下方法并未真正递归地考察所有可能的子问题。由于没有频繁的递归函数调用的开销，自底向上方法的时间复杂度函数通常具有更小的系数
 */

/**
	钢条切割
	钢条切割问题还存在一种相似的但更为简单的递归求解方法：将钢条从左边切割下长度为i的一段，只对右边剩下的长度为n-i的一段继续进行切割，对左边的一段则不再进行切割。这样得到的公式为：。这样原问题的最优解只包含一个相关子问题（右端剩余部分）的解，而不是两个。
 */

/**
	钢条切割问题---自顶向下 带备忘录略
	自底向上方法
	@n 钢条长度
	@p map 价值
 */
 func  BottomUpCutRod(n int,p map[int]int)int {
	//声明一个最优解集
	var dp []int
	dp = append(dp,0)  //dp[0] = 0

	//dp[i] = max(r[i - j],p[j]) 1<=j <= i
	for i := 1; i <= n ; i++ {
		q := -1
		for j := 1; j <= i; j ++ {
			temp := p[j] + dp[i - j]
			if q < temp {
				q = temp
			}
		}
		dp = append(dp,q)
	}

	//返回
	return dp[n]
 }

 func main() {
 	q := make(map[int]int)
 	q[1] = 1
 	q[2] = 5
 	q[3] = 8
 	q[4] = 9
 	q[5] = 10
 	q[6] = 17
 	q[7] = 17
 	q[8] = 20

 	fmt.Println(BottomUpCutRod(7,q))

 	var change []int
 	change = append(change,1)
 	change = append(change,2)
 	change = append(change,5)
 	change = append(change,21)
 	change = append(change,25)

 	fmt.Println(GiveChange(105,change))

 	jumpArr := []int{3,2,1,1,4}
 	fmt.Println(canJump(jumpArr))
 }

 /**
 	金币找零动态规划
  */
  func GiveChange(n int,change []int) int{
  	//声明保存最优解
  	var dp []int
  	dp = append(dp,0)  //金币为0时为0

  	//找零
  	for i := 1; i <= n; i++ {
  		//寻找此时决策
  		q := n + 1
  		for _,v := range change {
  			//判断
  			if v > i {
  				break
			}
  			temp := 1 + dp[i - v]

  			if q > temp {
  				q = temp
			}

		}
  		dp = append(dp,q)
	}



  	//返回
  	return dp[n]

  }
  /**
  给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
   */
func canJump(nums []int) bool {
	var dp []int
	dp = append(dp,nums[0])
	for i := 1; i < len(nums); i ++ {
		if dp[i - 1] < i {
			return false
		}
		dp  = append(dp, maxJump(dp[i - 1],nums[i] + i))
	}
	return true

}
func maxJump(a,b int) int {
	if  a > b {
		return a
	} else {
		return b
	}
}

