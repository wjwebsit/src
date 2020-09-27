package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

/**
	接口值：概念上讲一个接口值由两部分---接口值和接口类型。他们被称为接口的动态类型和动态值。
	对于go这种静态类型的语言，类型是编译期的概念因此一个类型不是一个值。
 */



func main() {
	//用io.Write接口来举例
	var w io.Writer //在go语言中变量总是被第一个明确的值初始化。接口也不例外
	//此时w为一个接口的零值 即 w的type和value 都为nil,一个接口值基于它的动态类型被描述为空或非空，所以这是一个空的接口值
	//接口可以通过nil来判断是否为空值
	fmt.Println(w == nil)  //true

	//空值调用失败
	//w.Write([]byte{'h','e','l','l','o'}) // invalid memory address or nil pointer dereference

	//w.type = os.*File os.value = os.File fd=1
	w = os.Stdout
	fmt.Println(w.Write([]byte{'h','e','l','l','o'})) // hello
	fmt.Println(reflect.TypeOf(w)) //*os.File
	fmt.Println(reflect.ValueOf(w).Type()) //*os.File

	w = new(bytes.Buffer) //赋值借口需要指针
	fmt.Println(w.Write([]byte("hello"))) //5,nil
	fmt.Println(reflect.TypeOf(w).String()) //*bytes.Buffer
	fmt.Println(reflect.ValueOf(w)) //hello

	//一个包涵nil指针的不是nil接口---就是一个接口被赋值成了nil它就不能在调用这个接口的方法了
	var f interface{} //空接口
	w  = nil
	f = w
	fmt.Println(f.(int)) // interface {} is nil, not int
}