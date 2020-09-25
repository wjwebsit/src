package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	算法描述:
	给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:
输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]

 */
func restoreIpAddresses(s string) []string {
	//声明返沪值
	var ipArr []string

	//获取数组长度
	length := len(s)

	//判断
	if length < 4 || length > 12 {
		return  ipArr
	}

	//判断短点次数--标准ip的最大长度即每位均有3为
	cutSum := 12 - length

	//回溯解决
	restoreIpAddressHS(s,0,cutSum,0,[]string{},&ipArr)

	//返回
	return ipArr


}
func restoreIpAddressHS (s string,start int,cutSum int,cur int,tmpResult []string,result *[]string) bool {
	//判断终止
	if len(tmpResult) > 4 {
		return  false
	}

	//判断是否完成
	if start >= len(s) {
		//判断
		if len(tmpResult) == 4{//ip完成
			//写入
			*result = append(*result,strings.Join(tmpResult,"."))

			return true
		} else {//未分割完
			return false
		}
	}


	//分割
	if  cur == cutSum { //当前无法分割
		//三位取
		if s[start] == '0'  { //不合法
			return false
		}

		//截取
		var str string
		if start + 3 > len(s) {
			str =  string(s[start:])
		} else {
			str = string(s[start : start+3])
		}

		//转换为整数---超过255不合法
		if num,_ := strconv.Atoi(str); num > 255 {
			return false
		}

		//写入
		tmpResult = append(tmpResult,str)

		//DFS
		res := restoreIpAddressHS(s,start + 3,cutSum,cur,tmpResult,result)

		//回溯
		tmpResult = tmpResult[0:len(tmpResult) - 1]
		return  res

	} else {//当前可以分割和不分割
		//三位取
		if s[start] == '0'  { //不合法---只能分割0
			//写入
			tmpResult = append(tmpResult,"0")

			//DFS
			res :=  restoreIpAddressHS(s,start + 1,cutSum,cur + 1,tmpResult,result)

			//回溯
			tmpResult = tmpResult[0:len(tmpResult) - 1]
			return res


		} else {
			//截取
			var str string
			if start + 3 > len(s) {
				str =  string(s[start:])
			} else {
				str = string(s[start : start+3])
			}

			//转换为整数---超过255不合法
			if num,_ := strconv.Atoi(str); num > 255 {
				//只能截取
				tmpResult = append(tmpResult,string(s[start: start + 1]))

				//DFS
				res1 := restoreIpAddressHS(s,start + 1,cutSum,cur + 1,tmpResult,result)

				//回溯
				tmpResult[len(tmpResult) - 1] =  string(s[start: start + 2])
				res2 := restoreIpAddressHS(s,start + 2,cutSum,cur + 1,tmpResult,result)
				tmpResult = tmpResult[0:len(tmpResult) - 1]

				//返回
				return res1 || res2

			} else {
				//当前不分割 111
				tmpResult = append(tmpResult,str)
				res1 := restoreIpAddressHS(s,start + 3,cutSum,cur,tmpResult,result)

				//分割1
				var res2 = false
				if str != string(s[start:start + 1]) {
					tmpResult[len(tmpResult)-1] = string(s[start : start+1])
					res2 = restoreIpAddressHS(s, start+1, cutSum, cur+1, tmpResult, result)
				}

				//分割 11
				var res3 = false
				if start + 2 < len(s) && string(s[start : start+2]) != str {
					tmpResult[len(tmpResult)-1] = string(s[start : start+2])
					res3 = restoreIpAddressHS(s, start+2, cutSum, cur+1, tmpResult, result)
				}


				//回溯
				tmpResult = tmpResult[0:len(tmpResult) - 1]

				return res1 || res2 || res3
			}

		}

	}
}
func main() {
	ip := "1111"
	fmt.Println(restoreIpAddresses(ip))
}