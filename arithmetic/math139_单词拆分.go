package main

import "strings"

/**
	题目描述:
	给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
 */
/**
   采用回溯算法并条件缓存
*/
func wordBreak(s string, wordDict []string) bool {
	//声明缓存
	cache := map[string]bool{
		"" : true,
	}

	//返回
	return wordBreakHS(s,wordDict,&cache)
}
/**
	*@param s string 字符串
	*@param wordDict []string 字典
	*@param cache *map[string]bool 缓存
	*@param isOk *bool 结束标识
	@return
 */
func wordBreakHS(s string ,wordDict []string,cache *map[string]bool) bool{
	//判断缓存
	if _,exist := (*cache)[s];exist {
		return (*cache)[s]
	}

	//遍历决策
	for _,word := range wordDict {
		//获取首次出现的位置
		index := strings.Index(s,word)

		//如果字符串不存在
		if index == -1 {
			continue
		}

		//判断0-index -1 && index + len(v) :
		left := wordBreakHS(string(s[:index]),wordDict,cache)
		if left == false {
			(*cache)[s] = false
			continue
		}

		(*cache)[s] = wordBreakHS(string(s[index + len(word):]),wordDict,cache)

		//判断
		if (*cache)[s] == true {
			break
		}

	}

	//返回
	return (*cache)[s]
}
