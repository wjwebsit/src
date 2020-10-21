package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	算法描述
	给定一个正整数数组 w ，其中 w[i] 代表下标 i 的权重（下标从 0 开始），请写一个函数 pickIndex ，它可以随机地获取下标 i，选取下标 i 的概率与 w[i] 成正比。

例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。

也就是说，选取下标 i 的概率为 w[i] / sum(w) 。

 

示例 1：

输入：
["Solution","pickIndex"]
[[[1]],[]]
输出：
[null,0]
解释：
Solution solution = new Solution([1]);
solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。
示例 2：

输入：
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
输出：
[null,1,1,1,1,0]
解释：
Solution solution = new Solution([1, 3]);
solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。

由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
[null,1,1,1,1,0]
[null,1,1,1,1,1]
[null,1,1,1,0,0]
[null,1,1,1,0,1]
[null,1,0,1,0,0]
......
诸若此类。
 

提示：

1 <= w.length <= 10000
1 <= w[i] <= 10^5
pickIndex 将被调用不超过 10000 次

 */
type Solution struct {
	sum []int
}


func Constructor3(w []int) Solution {
	var s Solution
	s.sum = make([]int,len(w))
	//计算抽奖概率
	sum := 0
	for i := 0; i < len(w);i ++ {
		sum += w[i]
		s.sum[i] = sum
	}

	//返回
	return s
}
/**
	根据权重随机返回
 */
func (this *Solution) PickIndex() int {
	//生成随机数
	rand := this.randInt(1,this.sum[len(this.sum) - 1])

	//二分查找最右边的
	s,e := 0,len(this.sum) - 1
	for s <= e {
		mid := s + (e - s) >> 1
		if this.sum[mid] >= rand {
			e = mid - 1
		} else {
			s = mid + 1
		}

	}

	//返回
	return s

}
/**
随机数
*/
func (this *Solution) randInt(min, max int) int{
	//设置
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min + 1)
}
func main() {
	w := []int{1,3}
	s := Constructor3(w)
	fmt.Println(s.PickIndex())
}