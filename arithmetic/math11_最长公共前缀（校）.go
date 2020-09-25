package main

import "fmt"

/**
	算法描述：
	编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
 */

 func main () {
	 //mp := map[string][]string{}
	 strs := []string{"flower","flow","flight"}
	 fmt.Print(longestCommonPrefix(strs))

 }

 /**
 思路采用折半查找的策略
  */
func longestCommonPrefix(strS []string) string {
	//判断临界
	if len(strS) == 0 {
		return ""
	}
	if len(strS) == 1 {
		return strS[0]
	}
	//定义maxPre 无穷 这里取-1就行了
	maxPre := -1

	//定义map
	mp := map[int][]string{}


	//原数组22比较找到最大前缀
	i:= 0

	for i < len(strS) {
		//后者
		j := i + 1 //二者构成比较数组

		//判断是否为临界 如果临界最后一个和前一个比较
		if j == len(strS) {
			j = i - 1
		}

		//找到22比较的局部最大前缀--受最小前缀的约束
		temPre := getPreSuffix(strS[i],strS[j])

		if temPre != ""{
			if maxPre == -1 {
				maxPre = len(temPre)
				mp[maxPre] = append(mp[maxPre],temPre)

			} else if (maxPre >= len(temPre)) {
				if findSliceIndex(temPre,mp[maxPre]) == -1 {
					maxPre = len(temPre)
					mp[maxPre] = append(mp[maxPre], temPre)
				}
			}

		} else {//出现了不存在的情况直接结束了
			maxPre = -1
			break
		}
		i+= 2
	}

	//判断
	if maxPre == -1 {
		return ""
	} else { //折半查找
		return longestCommonPrefix(mp[maxPre])
	}

}

/**
	寻找公共的最大前缀
 */
func getPreSuffix(s1,s2 string)string {
	i:= 0
	maxPre := ""

	for i < len(s1) && i < len(s2) {
		if s1[i] == s2[i] {
			maxPre += string(s1[i])
			i++
		} else {
			break
		}

	}
	//返回
	return maxPre
}
/**
查询在元素在切片的位置
 */
func findSliceIndex(key string,s []string) int{

	for i,value := range s {
		if value == key {
			return i
			break
		}
	}

	return -1
}

