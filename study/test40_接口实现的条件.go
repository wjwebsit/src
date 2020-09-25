package main

import (
	"fmt"
)

/**
	一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。
 */

/**
	接口中只能有方法不能有变量
 */
type animal interface {
	say()string
	eat()string
}
type cAt interface { //猫科动物
	isSleepDay()bool
}

type cat struct {
	name string
}
func (c *cat) say() string {
	return "miao"
}
func (c *cat) eat()string {
	return "fish"
}

func (c *cat)isSleepDay()bool {
	return true
}
func construct(name string) cat {
	var c cat
	c.name = name
	return c
}

func main() {
	//实力化cat
	c := construct("cat")
	fmt.Println(eat(&c))

	//主动调用
	var animal animal
	animal = &c
	fmt.Println(animal.eat())

	//空借口 任何变量都实现了空借口---interface{}
	var number interface{}
	number = 1
	number = 1.0
	number = float64(344)

	//获取值--变量名（.T） 前提是interface 类型
	fmt.Println(number.(float64) + 2.0) //346

	//1个类型可以有多中类型
	fmt.Println(isCat(&c))
}
/**
	利用接口来约束参数--Eat只能传递animal
 */
func eat(a animal)string {
	return a.say()
}

func isCat(c cAt) bool {
	return c.isSleepDay()
}

