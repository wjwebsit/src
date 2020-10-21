package main

import "fmt"

/**
	算法描述：
	给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。


示例 1：
输入：["bella","label","roller"]
输出：["e","l","l"]
示例 2：
输入：["cool","lock","cook"]
输出：["c","o"]
 
提示：
1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] 是小写字母
 */
func commonChars(A []string) []string {
	//声明返回值
	var result []string
	if len(A) == 0 {
		return result
	}
	//统计每个字符串中出现的次数

	mHash := make(map[byte]int)
	for j := 0;j < len(A[0]);j++ {
		mHash[A[0][j]] += 1
	}
	for i := 0; i < len(A);i++ {
		//存在空的字符串
		if len(A[i]) == 0 {
			return result
		}
		hash := make(map[byte]int)
		for j := 0;j < len(A[i]);j++ {
			hash[A[i][j]] += 1
		}
		//直接终止
		if len(mHash) == 0 {
			break
		}

		for k,v := range mHash {
			if _,ok := hash[k];ok {//存在
				mHash[k] = myMin(v,hash[k])
			} else {//不存在
				delete(mHash,k)
			}
		}
	}

	//构造返回结果
	if len(mHash) == 0 {
		return result
	}
	for k,v := range mHash {
		for i := 1;i <= v; i ++ {
			result = append(result,string(k))
		}
	}
	return result

}
func myMin(a,b int)int {
	if a > b {
		return b
	}
	return a
}

func main() {
	str := []string{"bella","e"}

	fmt.Println(commonChars(str))
}