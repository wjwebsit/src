package main

import "fmt"

/**
	数据类型转换
	T()  int 与 float64可以自由转换  *float64转int 精度丢失 会去掉小数部分  *float32与float64 go语言根据运行环境来判断而非数值大小
	bool 类型无法和其他类型转化
	int(),float64() 无法将字符串转化
	string() 可以转化byte,rune 类型 本质是将int的值当作字符集转化成对应的字符并成为字符串类型
*/

func main() {
	var a int = 90
	b := 90.9 //运行环境默认为float32或float64 --一般为float64

	//avg := (float32(a) + b) /2  **b实际为float64
	avg := (float64(a) + b) /2
	fmt.Println(avg)
	fmt.Print(int(b)) //返回为90 精度丢失

	/*fg := true
	int(fg)
	bool(b)*/  //bool类型无法转化

	c := string(a) //float64无法转化成字符串
	fmt.Printf("%s",c) //90 ==Z

	var d byte = 'Z'
	d += 7
	fmt.Printf(string(d)) //原理同c变量,输出为a

}
