package main

import "fmt"

/**
	算法描述:
	所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。

编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。

 

示例：

输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC", "CCCCCAAAAA"]

 */
func findRepeatedDnaSequences(s string) []string {
	//声明返回值
	var result []string

	//判断
	if len(s) <= 10 {
		return result
	}

	//定义map
	var mp  = make(map[string]int)
	var cache = make(map[string]int)

	//滑动窗口
	for i := 0; i + 10 <= len(s); i++ {
		str := string(s[i: i + 10])
		mp[str] ++
		if mp[str] > 1 && cache[str] == 0 {
			result = append(result,str)
			cache[str] = 1
		}
	}
	//返回
	return  result
}
func main()  {
	s := "AAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}