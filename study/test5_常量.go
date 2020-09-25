package main

import "fmt"

/**
	常量的值无法被修改
	常量的数据类型 只能为简单数据类型 string int float bool
	const name  T = 表达式  名称一般全为大写 如果权限首字母小写 其余大写
	const name = 表达式 隐示
	const (
		name1 =
		name2 =
	) 常量组形式
 */

func main() {
	const NAME = "赵宁"
	//NAME = "xuxiao"常量无法修改
	const A byte = 'a'
	fmt.Print(A)
	const (
		UNKOWN = iota
		FEMALE = iota
		MAN = iota
	) //常量组 一般用作枚举

	//fmt.Print(iota)iota只能用在常量组中
	fmt.Print(MAN) //MAN前面有2个 故为2 只是统计const组

}