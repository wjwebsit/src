package main
/**
	算法描述
	二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字不在0和1之间，或者无法精确地用32位以内的二进制表示，则打印“ERROR”。

示例1:

 输入：0.625
 输出："0.101"
示例2:

 输入：0.1
 输出："ERROR"
 提示：0.1无法被二进制准确表示
提示：

32位包括输出中的"0."这两位。
 */
func printBin(num float64) string {
	//判断
	if num <= 0 || num >= 1 {
		return "ERROR"
	}

	//转化---判断长度
	carry := 0.5

	//定义返回
	str := "0."

	for len(str) < 32 && num > 0 {
		//判断
		if num >= carry {
			str += "1"
			num -= carry
		} else {
			str += "0"
		}

		carry /= 2

	}

	//返回
	if num != 0 {
		return "ERROR"
	}

	return str


}