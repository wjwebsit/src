package main

import (
	"fmt"
	"strconv"
)

/**
算法描述
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。



示例 1:

输入: 1
输出: "1"
示例 2:

输入: 4
输出: "1211"
 */

func countAndSay(n int) string {
	//判断
	if n < 0 {
		return ""
	}

	//构造报数数组n=1 情况
	mp := map[int][]int{}
	mp[1] = []int{1}

	//从1开始构造
	for i := 2; i <=n ; i ++ {
		mp[i] = makeAndSay(mp[i - 1])


	}

	//处理结果
	rtStr := ""
	for _,v := range mp[n] {
		rtStr += strconv.Itoa(v)
	}

	//返回
	return rtStr
}

/**
	构造数组
 */
func makeAndSay(arr []int) []int {
	//定义返回值
	rtResult := []int{}

	//循环
	for i := 0; i < len(arr); {
		j := i + 1

		//寻找重复
		for j < len(arr) && arr[j] == arr[i] {
			j ++
		}


		//读取报文
		rtResult = append(rtResult,j - i )
		rtResult = append(rtResult,arr[i])

		//j赋值给i
		i = j
	}

	//返回
	return rtResult

}

func main() {
	fmt.Println(countAndSay(4))
}


