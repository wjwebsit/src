package main

import (
	"fmt"
)

/**
	题目描述:
	给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

说明：

分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
示例 2：

输入:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
输出:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
解释: 注意你可以重复使用字典中的单词。
示例 3：

输入:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
输出:
[]
 */
 /**
 	利用回溯算法加缓存机制
  */
func wordBreakII(s string, wordDict []string) []string {
	//声明结果集缓存 -result["catsand"] = ["cat sand","cats and"]
	result := map[string][]string{}

	//构造map
	wordMap := map[string]int{}

	//循环
	for _,v := range wordDict {
		wordMap[v] = 1
	}

	//回溯算法求解
	return wordBreakIIHS(s,wordMap,&result)


}

func wordBreakIIHS(s string,wordDict map[string]int,result *map[string][]string) []string {
	//判断缓存
	if _,ok := (*result)[s]; ok {//表示缓存存在

		return (*result)[s]
	}
	//声明
	var res []string

	//分割
	for i := 1; i <= len(s); i ++ {
		str := string(s[:i])
		if wordDict[str] == 1 {
			//判断是否包含
			rightRes := wordBreakIIHS(string(s[i:]),wordDict,result)

			//如果不包含
			if len(rightRes) > 0 {
				for _,v := range rightRes {
					res = append(res,str + " " + v)
				}
			} else if i == len(s) {
				res = append(res,str)
			}
		}

	}
	(*result)[s] = res
	//返回
	return (*result)[s]

}
func main() {

	s := "leetcode"
	wordDict := []string{"leet","code"}

	result := wordBreakII(s,wordDict)
	mp := map[string]int{}

	//消除冗余
	for _,v := range result {
		if mp[v] == 0 {
			mp[v] = 1
		}
	}
	fmt.Println(result)
}