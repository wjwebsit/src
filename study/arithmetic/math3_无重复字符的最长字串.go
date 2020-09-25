package main

import "fmt"

//算法题目：
/**
给定一个字符串，你请找出其中不含有重复字符的  最长子串 的长度。
输入： “abcabcbb”
 输出： 3
 解释：因为无重复字符的最长子串是"abc"，所以其长度为3。
 */
 func main() {
 	str1 := "abcaebcbb"
	fmt.Print(myMethod(str1))

 }

 //暴力破解法---
 func myMethod(str string) int {

 	 max := 0
 	//起始索引
 	for begin := 0; begin < len(str);begin++ {
 		//获取起始元素
 		var sValue byte = str[begin]

 		//定义临时切片
 		temp := []byte{}
 		temp = append(temp,sValue)


 		//下一个节点索引
 		end := begin + 1

 		//判断下线不存在前面并且没有越界
 		for end < len(str) && findIndex(str[end],temp) == -1 {
			temp = append(temp,str[end])
 			end ++
		}

 		if len(temp) > max {
 			max = len(temp)
		}

	}

 	return max

 }
/**
查询在元素在切片的位置
 */
 func findIndex(key byte,s []byte) int{

 	for i,value := range s {
 		if value == key {
 			return i
 			break
		}
	}

 	return -1
 }
