package main

import "fmt"

/**
 算法描述：
将一个给定字符串根据给定的行数，以从上往下，从左到右进行Z字形排列。

比如输入字符串为"LEETCODEISHIRING" 行数为3时，排列如下：

LCIR
ETOESIIG
EDHN
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

示例：

输入： s =“LEETCODEISHIRING”，numRows = 4
 输出：  “LDREOEIIECIHNTSG”
 解释：

LDR
EOEII
ECIHN
TSG


 */

func main() {
	//测试
	//str := "LEETCODEISHIRING"
	//row := 3

	fmt.Println(myChange("AB",1))
}
/**
@param string s 进行转换的字符串
@param int row 转化的行数
@return string 重新生成的字符串
 */
func myChange(s string ,row int) string {
	if row == 1 {
		return s
	}

	//定义数组row存放
	indexRow := -1

	//定义方向
	dirction := 0

	//定义存放字符的二维切片
	var result [][]byte


	//遍历字符串
	for i := 0; i < len(s); i++ {
		//初始化
		if i < row {
			temp := []byte{}
			temp = append(temp,s[i]) //注意二维切片的初始化

			indexRow ++
			result = append(result,temp)


		} else {//初始化完成

			if indexRow == row - 1 {
				//方向回升
				dirction = 1
			}

			if indexRow == 0 {
				dirction = 0
			}

			//填充数据向上
			if dirction == 1 {
				indexRow --
				result[indexRow] = append(result[indexRow],s[i])
			}

			if dirction == 0 {
				indexRow ++
				result[indexRow] = append(result[indexRow],s[i])
			}

		}


	}

	//定义返回
	rtStr := ""

	for _,value := range result {
		rtStr += string(value[:])
	}

	//返回新的字符串
	return rtStr



}