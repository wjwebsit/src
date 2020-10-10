package main

import "fmt"

/**
	方法值和方法表达式
	方法值：就是将方法作为参数赋值
	方法表达式：就是参数是结构体，不同的结构体实例，来执行方法
 */

type point struct {
	x,y int
}
/**
	获取x
 */
func ( p *point) getX() int{
	//返回
	return p.x
}
/**
	获取y
 */
func (p point) getY() int {
	//返回
	return p.y
}
/**
	设置x
 */
func (p *point) setX(x int)  {
	p.x = x
}

/**
	设置y
 */
func (p *point) setY(y int) {
	p.y = y
}
/**
	闭包
 */
func (p *point) getArea(f1 func()int,f2 func()int) int {
	return  f1() * f2()
}

func main() {
	//方法值
	p1 := point{12,24}
	getX := p1.getX
	fmt.Println(getX()) //12

	//传递闭包
	fmt.Println(p1.getArea(p1.getX,p1.getY)) //288

	//方法表达式
	getY := (*point).getY
	//fmt.Println(getY(&p1)) //错误结构体参数为传引用(*point).getY---很关键
	fmt.Println(getY(&p1)) //24

	(*point).setX(&p1,20)
	fmt.Println(getX()) //20


}