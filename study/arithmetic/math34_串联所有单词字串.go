package main

import (
	"fmt"
	"strings"
)

/**
	算法描述
	给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words  中所有单词串联形成的子串的起始位置。

 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoor" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
 */
/**
	回溯算法
 */
func findSubstring(s string, words []string) []int {
	//定义返回结果
	var result []int

	//判断
	if len(s) == 0 || len(words) == 0 {
		return result
	}

	//开始遍历决策
	trySubstring(s,words,"",0,&result)

	//返回结果
	return result
}

/**
	@param str 字符串
	@param choice 决策数组
	@param search 上一步决策后的str
	@param n 选择的数量 当n == len(choice) 结束
	@param result 结果

 */
func trySubstring(str string,choice []string,search string,n int ,result *[]int) {
	//判断结束
	if n == len(choice) {
		//判断是否匹配
		if strings.Index(str,search) != -1  && !inResult(*result,strings.Index(str,search)){
			//记录结果
			*result = getIndexStr(str,search,*result)
		}
		//回溯
		return
	}
	//开始遍历决策--增加 历史
	mp := map[string]int{}

	//决策
	for i := 0; i < len(choice); i ++ {
		//判断是否选择过
		if choice[i] == "-" || mp[choice[i]] == 1{
			continue
		}

		//判断决策
		if strings.Index(str,search + choice[i]) == -1 {//当前决策失败
			continue

		} else { //下一次尝试 不需要choice[i] -- choice 置为"-"
			temp := choice[i]
			choice[i] = "-"
			trySubstring(str,choice,search + temp,n + 1,result)

			//恢复
			choice[i] = temp

			//写入历史
			mp[temp] = 1
		}
	}

}
/**
	判断是否在切片里
 */
func inResult(result []int,value int) bool {
	if len(result) == 0 {
		return false
	}
	for _,v := range  result {
		if v == value {
			return true
		}
	}
	return false
}
/**
	寻找a在str 出现的位置
 */
func getIndexStr (str string,a string,result []int)[]int {
	//str 指针
	i := 0

	//a 指针
	j := 0

	for ;i < len(str);i++ {
		//判断
		if len(str) - i < len(a) {
			break
		}
		//判断
		if str[i] == a[j] {
			for j < len(a) && str[i + j] == a [j] {
				j++
			}

			//判断是否匹配成功
			if j == len(a) &&  !inResult(result,i) {
				result = append(result,i)
			}
			j = 0
		}

	}


	//返回值
	return result

}

func main() {
	str :="barfoothefoobarman"
	words := []string{"foo","bar"}
	fmt.Println(findSubstring(str,words))
}
