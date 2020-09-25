package main

import "math"

/**
	题目描述:
	给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
 */

 /**
 	单词接龙
  */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//判断
	if len(wordList) == 0 {
		return 0
	}

	//将wordList转换为散列 来替代 in_array
	wordMp := map[string]int{}
	for _,v := range wordList {
		wordMp[v] = 1
	}

	//判断最终单词是否在结果中
	if wordMp[endWord] == 0 {
		return 0
	}

	//缓存mp来记录当前目标单词的索引是否转化过
	cache := map[int]int{}

	//初始化
	for i := 0; i < len(endWord);i++ {
		cache[i] = 0 //表示为转化过
	}

	//定义返回值
	min := math.MaxInt64

	//采用回溯算法


}
/**
	count为已经替换的次数
 */
func ladderLengthHS(beginWord,endWord string,cache map[int]int,wordMp map[string]int,count int,min *int,success int) {
	//判断终止条件
	if success == len(beginWord) {
		//比较
		if *min > success {
			*min = success
		}
		return
	}

	//先判断是否直接能转换


}
