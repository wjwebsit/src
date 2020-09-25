package main

import (
	"fmt"
	"strings"
)

/**
	算法描述
	稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

示例1:

 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
 输出：-1
 说明: 不存在返回-1。
示例2:

 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
 输出：4
提示:

words的长度在[1, 1000000]之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sparse-array-search-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findString(words []string, t string) int {
	//判断数组
	if len(words) == 0 {
		return  -1
	}
	s,e := 0,len(words) - 1

	for s <= e {
		mid := s + (e - s) / 2

		//判断是否为空格
		if words[mid] == "" {
			//[mid+1,e - 1] 之间寻找非空格
			mid ++
			for mid < e && words[mid] == "" {
				mid ++
			}

			//判断
			if mid == e {//mid和e之间全为空格
				e = mid - 1
				mid = s + (e - s) / 2
			}
		}
		res := strings.Compare(words[mid],t)
		//判断
		if res == 0 {
			return  mid
		} else if res > 0 {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return  -1


}

func compareStr(s1,s2 string) bool {

	return true
}
func main() {
	fmt.Println(strings.Compare("b","a"))
}