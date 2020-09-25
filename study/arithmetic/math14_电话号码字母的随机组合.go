package main

import "fmt"

/**
	算法描述
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

    给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
	示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 */

 func main() {
	fmt.Print(letterCombinations("23"))
 }
/**
字母的随机组合
 */
func letterCombinations(digits string) []string {
	//构造字典
	mp := map[byte][]string{}
	mp['2'] = []string{"a","b","c"}
	mp['3'] = []string{"d","e","f"}
	mp['4'] = []string{"g","h","i"}
	mp['5'] = []string{"j","k","l"}
	mp['6'] = []string{"m","n","o"}
	mp['7'] = []string{"p","q","r","s"}
	mp['8'] = []string{"t","u","v"}
	mp['9'] = []string{"w","x","y","z"}

	//声明result
	result := []string{}

	//进行遍历
	for i := 0; i < len(digits); i++ {
		//判断当前是否为1
		if digits[i] == '1' {
			continue
		}

		//判断是否为空
		if len(result) == 0 {
			result = mp[digits[i]]
			continue
		}

		//排列组合
		tempResult := []string{}
		for k := 0; k < len(result);k++ {
			for l := 0; l < len(mp[digits[i]]);l ++ {
				//组合
				tempResult = append(tempResult,result[k] + mp[digits[i]][l])
			}

		}

		//赋值result
		result = tempResult

	}

	//返回
	return result
}
