package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	算法描述:
	给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 */
func largestNumber(nums []int) string {
	//按高位进行排序--如果重复则在按次高位重复 注意 9 he 98 9>98
	var arr = make([][]string,10)

	//写入对相应arr
	for i := 0;i < len(nums);i++  {
		//转化为字符串
		num := strconv.Itoa(nums[i])
		arr[num[0]-'0'] = append(arr[num[0] - '0'],num)
	}


	//排序0不用排
	str := ""
	for i:= 9;i > 0;i-- {//降序
		//排序
		if len(arr) > 0 {
			str += strings.Join(mySortNum(arr[i]),"")
		}
	}
	if len(arr[0]) > 0 {
		if str != "" {
			str += strings.Join(mySortNum(arr[0]),"")
		} else {
			str = "0"
		}
	}

	//
	return str

}
func mySortNum(nums[]string) []string{
	//采用插入排序
	for i := 1;i < len(nums);i ++ {
		//获取当前key
		key := nums[i]

		//获取后一位
		j := i - 1
		for j >= 0 && compare(key,nums[j]) {
			//交换
			nums[j + 1] = nums[j]
			j --
		}
		nums[j + 1] = key
	}
	return nums

}

func compare(a,b string) bool {
	res1 := a + b
	res2 := b + a
	return res1 > res2
}
func main() {
	nums := []int{824,938,1399,5607,6973,5703,9609,4398,8247}
	fmt.Println(largestNumber(nums))
}