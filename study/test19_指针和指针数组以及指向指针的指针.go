package main

import "fmt"

/**
指针：指针是一个特殊的变量，它里面存储的数值被解释成为内存里的一个地址====变量赋值 无非是先找到变量的地址然后再地址上去做修改
A、指针的声明：
1、var 变量名 *T  此时指针实际上为nil  var p *int
2 、var 变量名 *T = &（后面为T类型的变量） var p *int = &a === p := &a
3、var 变量名 *T = new(T) 返回初始化T的地址 ===不是nil 值为T数据类型的默认值

B、指针的值访问
*指针变量 即为值

C、指针数组
var 名称 [length]*T ===即数组的每一项均为指针  [...]int{1，3} 这中长度不写明确的写法

D、指向指针的指针
var 名称 **T  即当前的变量指向的是一个指针的地址
访问值 **变量名


 */

func main() {
	/****A、指针的声明*****/
	//、var 变量名 *T ****不推荐 --没有实例化
  	var p1 *int //此时为nil
  	fmt.Println(p1)
  	//*p1 = 2  //错误因为p1此时为nil nil无效的地址


	//、var 变量名 *T = &（后面为T类型的变量） var p *int = &a === p := &a
	a := 1
	p2 := &a
	fmt.Println(p2)

	//、var 变量名 *T = new(T) 返回初始化T的地址 ---已经分配了内存实例化了
	p3 := new(int)
	fmt.Println(*p3) //0

	/******B、指针的值访问*********/
	fmt.Println(*p2) //1

	/******C、指针数组*********/
	//*p1 = 2
	var piArray  = [2]*int{p2,p3}
	for _,p := range piArray {
		fmt.Println(*p)
	}

	/*****D、指向指针的指针*****/

	b := 10
	pb := &b
	ppb := &pb

	fmt.Println(**ppb) //10

	/******E、go语言中函数体内如果传引用变量重新赋值则传引用失效*****/
	d := 1
	p4 := &d
	p5 := p4

	p5 = new(int)  //new 相当于是取的匿名变量的地址

	fmt.Println(*p4,*p5)

	E := 6

	test(&E)

	fmt.Print(E) //6 因为函数体内指针重新指向了一个新的地址


}
/**
测试指针的失效性
 */
func test(a *int) {
	a = new (int)
	*a = 2
}
