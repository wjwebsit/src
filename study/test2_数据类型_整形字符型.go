package main

import "fmt"

/**
go语言数据类型：
基本数据类型：整形、浮点型、布尔型、字符串、字符型
 */

func main() {
	//整形int
	/*
		无符号整形
		uint8 无符号小整形 0-(2^8-1) 间 取值为0-255 alias byte
		uint16 同上
		uint32 同上
		uint64 同上
		有符号整型
		int8 有符号 值域 -128-127
		int16
		int32  alias rune
		int64
		复合整形
		byte === uint8
		rune === int32
		uint === uint32||uint64 根据值的大小来动态分配
		int ==== int32 || int64 根据值的大小来动态分配
		uintptr 无符号整形 用于存放一个指针
	 */
	 //var a uint8 = 256 超出范围 overflows
	 var a byte = 255
	 fmt.Printf("%T",a)

	 var b rune = 455
	 fmt.Printf("%T",b)

	 /**
	 	字符型
	      Go语言字符有2种类型
	        1、byte型 代表ascii码的一个字符
	 		2、rune	型 代表UTF-8的一个字符 当处理中文时需要用到rune型
	  */
	  var bt byte = 'a'
	  fmt.Printf("%T %v",bt,bt) //返回a的ascii 的值97

	  var ru rune = 'a'
	  fmt.Printf("%T %v",ru,ru) //也可以实现

	  var rue rune = '许'
	  fmt.Printf("%T %v",rue,rue)

}
