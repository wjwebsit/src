package main

import "fmt"

/**
	算法描述：
	只猴子有100个香蕉，把香蕉搬到50米之外的家里，每次最多拿50个香蕉，每走一米要吃掉一个香蕉，最后家里能剩下几个香蕉
 */

/**
	*@param sum int 总的相交数目
	*@param distance 距离家的距离
	*@param get int 最大搬运量

 */
func banana(sum int,distance int,get int) int{
	//定义来回搬运的距离
	d := 0

	//来回搬的次数
	for sum >= (get + 2)  && d < distance{
		sum -= 3
		d ++
	}

	//判断是否可以来回搬到家
	if d == distance {
		return sum
	}

	//返回一次般的最后剩余
	last := sum - distance + d

	//若果为－ 返回0
	if last <= 0 {
		return 0
	}

	//返回结果
	return last

}

func main() {
	fmt.Println(banana(100,50,50))
}