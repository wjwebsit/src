package main

import "fmt"

//b := 3 这种简单的声明格式只能在函数域中声明,无法用于全局变量

var ( //--------批量声明变量方式 一般适用于全局变量的声明
	constA int //默认值为0
	constB float32 //默认值为0
	constC bool //默认值为 false
	constD func() bool //默认值为nil
	constE string //默认值为""
	constF int32 = 8 //支持表达式声明
)

func main() {
	//变量声明
	//1、var 变量名 类型
	var a int
	var p *int
	var arr []int //没有指定长度为切片
	var array [2]int //指定长度为数组，且数组中的每个元素默认值为数组类型的默认值，本案例int默认为0
	fmt.Println(a,p,arr,array)

	//2、var 变量名 类型 = 表达式(||函数)
	var b int = 1
	var constA *int = &b //变量名称无法重复 但可以与全局相同
	var str string = "许校"
	fmt.Println(b,*constA,str)

	//3、var 变量名 = 表达式 （弱类型的语法）---不推荐
	var c = 3.12
	var constB = new(float32)
	fmt.Println(c,*constB)

	//4、第3的简短写法 变量名 := 表达式  *变量 不能提前声明
	d := [4]int{1,5,3,4}
	pArr := &d
	fmt.Println(*pArr)

	//多变量赋值声明 var 变量1，变量2 = 值1，值2 ===== 变量1，变量2 := 值1，值2(推荐)
	var e1,e2 = 2,4.4
	fmt.Println(e1,e2)

	e3,e4 := 3,45
	fmt.Println(e3,e4)

	//交换a,b(ab类型相同) ===特性
	e3,e4 = e4,e3
	fmt.Println(e3,e4)

	//指针数组 [] *T
	var ptArr []*int
	ptArr = append(ptArr,p)
	fmt.Print(ptArr)

}
