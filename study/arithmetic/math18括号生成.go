package main

import "fmt"

/**
算法描述：
给出  Ñ  代表生成括号的对数，请你写出一个函数，能够使其生成所有可能的并且有效值的括号组合。

例如，给出  n = 3，生成结果为：

[
  “（（（）））”，
  “（（）（））”，
  “（（））（）”，
  “（）（（））”，
  “（）（）（）”
]
 */

func main() {
	fmt.Println(generateParenthesis(3))
}
/**
	括号生成 -- 回溯算法
 */
func generateParenthesis(n int) []string {
	//定义返回值
	var rtResult []string

	//判断
	if n == 0 {
		return rtResult
	}

	//判断1
	if n == 1 {
		rtResult = append(rtResult,"()")
		return rtResult
	}

	//决策
	choice := []string{"(",")"}

	//回溯解决问题
	tryMake(n,choice,1,0,"(",&rtResult)

	//返回
	return rtResult




}
/**
	回溯问题
 */
func tryMake(n int,choice []string,lNum int,rNum int,str string,result *[]string) {
	//判断是否生成玩
	if lNum == rNum && rNum == n{

		//写入值;
		*result = append(*result,str)
		return
	}

	//每一步的决策
	for _,v := range choice {
		//判断允许条件
		if v == "(" && lNum < n {
			tryMake(n,choice,lNum + 1,rNum,str+"(",result)
		} else if v == ")" && rNum < lNum && rNum  < n {
			tryMake(n,choice,lNum,rNum + 1,str+")",result)
		}
	}

}



