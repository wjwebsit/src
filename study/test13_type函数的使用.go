package main

import "fmt"

func main() {
	//函数传引用通过指针
	s1:= 0
	sum(100,&s1)
	fmt.Println(s1)

	//闭包函数一般调用方式
	s2 := 0
	myMethod(50,&s2,sum)
	fmt.Println(s2)

	//闭包函数的常用调用方式---动态灵活
	fmt.Println(xuxiao(1,2,add))
	fmt.Println(xuxiao(1,2,jian))

}

//定义一个函数1~n之间的和
func sum(n int, s *int) {
	if n >= 1 {
		*s += n
		sum(n - 1,s)
	}
}
//闭包方式1
func myMethod(n int,p *int,sum func(int,*int)) {
	sum(n,p)
}

//闭包方式2
func add(a,b int) int {
	return a + b
}

func jian(a,b int) int {
	return a - b
}
//自定义类型方式-----推荐
type funcName func(int,int) int

func xuxiao(a,b int,name funcName) int {
	return name(a,b)
}

