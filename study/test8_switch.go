package main

import "fmt"

/**
	switch 数值  和 case 后面的数值类型必须相同
 */

func main() {
	var grade float32 =  4.5

	switch fmt.Sprintf("%T",grade) {
	case "float64","float32": //****float32 ||float64
		fmt.Print("浮点型")
		break
	case "int":
		fmt.Print("整型")
		break
	default:
		fmt.Print("不合法的数据类型")

	}
}
