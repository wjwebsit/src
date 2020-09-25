package main
/**
无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。

示例1:

 输入：S = "qwe"
 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
示例2:

 输入：S = "ab"
 输出：["ab", "ba"]
提示:

字符都是英文字母。
字符串长度在[1, 9]之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-i-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func permutation(S string) []string {
	//声明结果集
	var result []string

	//判断
	if len(S) == 0 {
		return result
	}

	//统计出现的字符集
	hash := make(map[byte]int)

	//遍历
	for i := 0; i < len(S); i ++ {
		hash[S[i]] ++
	}

	//
	dfsPermutation(hash,0,len(S),[]byte{},&result)

	return result
}


func dfsPermutation(hash map[byte]int,cur , sum int,str []byte,result *[]string) {
	//判断
	if cur == sum {
		*result = append(*result,string(str))
		return
	}

	for k,v := range hash {
		if v < 1 {
			continue
		}
		hash[k] --
		str = append(str,k)
		dfsPermutation(hash,cur + 1,sum,str,result)

		//回溯
		hash[k] ++
		str = str[:len(str) - 1]
	}

}