package main

import "fmt"

/**
	算法描述:
	给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...
示例 1:

输入: 1
输出: "A"
示例 2:

输入: 28
输出: "AB"
示例 3:

输入: 701
输出: "ZY"
 */
func convertToTitle(n int) string {
	//声明map
	mp := map[int]byte{1:'A',2:'B',3:'C',4:'D',5:'E',6:'F',7:'G',8:'H',9:'I',10:'J',11:'K',12:'L',13:'M',14:'N',15:'O',16:'P',17:'Q', 18:'R',19:'S',20:'T',21:'U',22:'V',23:'W',24:'X',25:'Y',26:'Z',}

	//判断
	if n <= 26 {
		return string(mp[n])
	}

	//返回值
	result := ""
	for n > 26 {
		if n % 26 == 0 {
			result = string(mp[26]) + result
			n = n / 26 - 1
		} else {
			result = string(mp[n % 26]) + result
			n = n / 26
		}

	}

	//最后的
	if n  == 0 {
		return  result
	} else {
		return  string(mp[n]) +  result
	}
}
func main()  {
	fmt.Println(convertToTitle(701))
}