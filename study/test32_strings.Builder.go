package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串构建器
	var builder strings.Builder

	//---插入一个字符
	builder.WriteByte('{')

	//---插入字符数组
	builder.Write([]byte{'h','e','l','l','o'})

	//获取构建器的长度--字符的个数
	fmt.Println(builder.Len()) // 6

	//插入字符串
	builder.WriteString("has")
	fmt.Println(builder.Len())  //6 + 3

	//扩容
	builder.Grow(12)

	//获取容量16 会预先分配内存
	fmt.Println(builder.Cap())

	//输出字符串
	fmt.Println(builder.String())

	//重置清空
	builder.Reset()
	fmt.Println(builder.String())
}
