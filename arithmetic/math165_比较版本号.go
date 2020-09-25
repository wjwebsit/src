package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	算法描述:
	比较两个版本号 version1 和 version2。
如果 version1 > version2 返回 1，如果 version1 < version2 返回 -1， 除此之外返回 0。

你可以假设版本字符串非空，并且只包含数字和 . 字符。

 . 字符不代表小数点，而是用于分隔数字序列。

例如，2.5 不是“两个半”，也不是“差一半到三”，而是第二版中的第五个小版本。

你可以假设版本号的每一级的默认修订版号为 0。例如，版本号 3.4 的第一级（大版本）和第二级（小版本）修订号分别为 3 和 4。其第三级和第四级修订号均为 0。
 

示例 1:

输入: version1 = "0.1", version2 = "1.1"
输出: -1
示例 2:

输入: version1 = "1.0.1", version2 = "1"
输出: 1
示例 3:

输入: version1 = "7.5.2.4", version2 = "7.5.3"
输出: -1
示例 4：

输入：version1 = "1.01", version2 = "1.001"
输出：0
解释：忽略前导零，“01” 和 “001” 表示相同的数字 “1”。
示例 5：

输入：version1 = "1.0", version2 = "1.0.0"
输出：0
解释：version1 没有第三级修订号，这意味着它的第三级修订号默认为 “0”。
 

提示：
版本字符串由以点 （.） 分隔的数字字符串组成。这个数字字符串可能有前导零。
版本字符串不以点开始或结束，并且其中不会有两个连续的点。
 */
func compareVersion(version1 string, version2 string) int {
	//分割字符串
	versionArr1 := strings.Split(version1,".")
	versionArr2 := strings.Split(version2,".")

	//循环比较
	var i int
	for i = 0; i < len(versionArr1) && i < len(versionArr2); i ++ {
		//依次比较，
		res := myCompareVersion(versionArr1[i],versionArr2[i])

		//判断
		if res != 0 {
			return  res
		}

	}

	//判断最后
	if i == len(versionArr1)  && i == len(versionArr2) {
		return 0
	}
	//判断
	if i < len(versionArr1) {
		//遍历
		for i < len(versionArr1) {
			res := myCompareVersion(versionArr1[i],"0")

			//判断
			if res != 0 {
				return res
			}
			i++

		}
		return 0

	} else {
		//遍历
		for i < len(versionArr2) {
			res := myCompareVersion("0",versionArr2[i])

			//判断
			if res != 0 {
				return res
			}
			i++

		}
		return 0
	}

}

func myCompareVersion(s1,s2 string) int {

	//去掉前导0
	var i ,j int
	for i = 0;i < len(s1); i ++ {
		if s1[i] != '0' {
			break
		}
	}
	for j = 0; j < len(s2); j ++ {
		if s2[j] != '0' {
			break
		}
	}


	//判断是否越界
	if i == len(s1) && j == len(s2) {
		return 0
	}

	//判断其它
	if i == len(s1) && j < len(s2) {//0 com 01
		return -1
	}
	if j == len(s2) && i < len(s1) {//01 com 000
		return  1
	}

	//均存在---去掉前导0
	s1 = string(s1[i: len(s1)])
	s2 = string(s2[j:len(s2)])


	//转为整数
	num1,_ := strconv.Atoi(s1)
	num2,_ := strconv.Atoi(s2)
	if num2 == num1 {
		return  0
	}
	if num1 > num2 {
		return  1
	}
	return  -1
}
func main() {
	fmt.Println(compareVersion("1.2","1.10"))
}