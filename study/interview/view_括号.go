package main
/**
	算法描述
	括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。

说明：解集不能包含重复的子集。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

 */
func generateParenthesis(n int) []string {
	//声明返回
	var result []string

	//判断
	if n == 0 {
		return result
	}
	makeGenerateParenthes(1,0,n,"(",&result)

	return result

}

func makeGenerateParenthes(left,right,n int,str string,result *[]string) {
	//判断
	if left == n && right == n {
		//写入结果
		*result = append(*result,str)
		return
	}

	//添加左括号
	if left < n {
		makeGenerateParenthes(left + 1,right,n,str + "(",result)
	}

	//添加右括号
	if right < left {
		makeGenerateParenthes(left,right + 1,n,str + ")",result)
	}


}