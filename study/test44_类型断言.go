package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/**
	类型断言是一个使用在接口值上的操作。语法看上去像x.(T)被称为类型断言。
	一个类型断言检查它操作对象的动态类型是否和断言的类型匹配
 */

func main() {
	//设置
	var w io.Writer //接口
	w = os.Stdout
	w.Write([]byte("hello"))

	//通过接口来获取对象
	b := make([]byte,5)
	f,_ := w.(*os.File)
	f.Read(b)
	fmt.Println(string(b))

	/*m := w.(*bytes.Buffer)  //如果直接x := w.(T) 相当于断言成功直接赋值给x ,但如果断言失败运行会出错
	fmt.Println(m)*/ // io.Writer is *os.File, not *bytes.Buffer

	//正规写法类型断言--判断
	if _,ok := w.(*bytes.Buffer);!ok {
		fmt.Println(ok)
	}
	fmt.Println("endof")

	//######接口和对象之间转换
	//1、对象赋值接口   接口名 = 对象指针（空接口类型除外）--如果对象不属于次接口 会报语法错误
	//2 接口还原对象  =x.(T) ,简易写法但是如果断言失败会错误
	//3、更好的类型断言 if f,ok := x.(T);ok {断言成功！，失败不会报错}
	var number interface{}
	number = 12
	//fmt.Println(number + 12) //此时interface类型无法做加法
	fmt.Println(number.(int) + 12) //转化成int就可以了
	number = 23.0
	fmt.Println(number.(float64) + 23) //46
	number = "123"
	fmt.Println(number.(string)) //123

}