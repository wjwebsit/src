package main

import "fmt"

/**
	算法描述
	给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
示例 2:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false
 */

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1) + len(s2) {
		return false
	}
	if s1 == "" {
		return  s3 == s2
	}
	if s2 == "" {
		return  s3 == s1
	}
	//判断
	if string(s3[:len(s1)]) == s1 || string(s3[:len(s2)]) == s2 {
		return false
	}
	cache := make(map[int]bool)
	return  isInterleaveHS(s1,s2,s3,0,0,0,cache)
}

func isInterleaveHS(s1,s2,s3 string,i,j,k int,cache map[int]bool) bool {
	//判断结束
	if k == len(s3) {
		return true
	}

	//决策
	isS1 := false
	isS2 := false

	if i < len(s1) && s3[k] == s1[i] {
		isS1 = true
	}
	if j < len(s2) && s3[k] == s2[j] {
		isS2 = true
	}

	if !isS1 && !isS2 {
		//写入缓存
		cache[k] = false

		return  false
	}

	//判断
	if isS1 == true && isS2 == false {
		cache[k] = isInterleaveHS(s1,s2,s3,i + 1,j,k + 1,cache)

	} else if isS2 == true && !isS1 {
		cache[k] = isInterleaveHS(s1,s2,s3,i,j + 1,k + 1,cache)

	}  else {
		cache[k] = isInterleaveHS(s1,s2,s3,i + 1,j,k + 1,cache) || isInterleaveHS(s1,s2,s3,i,j + 1,k + 1,cache)
	}

	//返回
	return cache[k]

}
func main() {
	s1:= "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	fmt.Println(isInterleave(s1,s2,s3))
}
