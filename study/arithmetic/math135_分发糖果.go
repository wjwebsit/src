package main

import "fmt"

/**
	算法描述:
	老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:

输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例 2:

输入: [1,2,2]
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
*/
func candy(ratings []int) int {
	//判断
	if len(ratings) == 0 {
		return 0
	}

	//定义总需的数目
	sum := 0

	//声明分发的数组
	var candy []int
	candy = append(candy,1) //candy[0] = 1

	for i:= 1;i < len(ratings); i++ {
		//判断是否比前一个优秀
		if ratings[i] > ratings[i - 1] {//比前一个优秀
			candy = append(candy,candy[i - 1] + 1)

		} else if ratings[i] < ratings[i - 1]{//不比前一个优秀
			//当前分一个
			candy = append(candy,1)

			//更新前面的 满足属性
			j := i
			for (j-1) >= 0 && ratings[j-1] > ratings[j] && candy[j - 1] == candy[j] {
				candy[j-1] ++
				j --

			}


		} else {//和前一个相等---发一个就可以了
			candy = append(candy,1)
		}
	}

	//统计
	for i := 0; i < len(candy);i++ {
		sum += candy[i]
	}
	//返回
	return sum
}

func main() {
	fmt.Println(candy([]int{1,3,2,2,1}))
}


