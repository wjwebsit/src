package main

import (
	"fmt"
	"io"
	"strings"
)

/**
	io接口是golang中实现最多的接口
	io.Reader 表示一个读取器将数据读到某一个缓冲区。在缓冲区中数据可以被流式的传输和使用
	type Reader interface { ---结构体
		Read(p []byte) (n int, err error)
	}
 */

func main() {
	var reader = strings.NewReader("www.hello,word")

	//定义读取buffer---
	fmt.Println(reader.Len()) //返回未读取的字符长长度
	fmt.Println(reader.Size()) //返回要读取字符长的长度固定不变

	var content []byte
	/*for {
		//读取的字节流大小---每次读取这么多
		var buffer = make([]byte,9)

		//读取资源到
		n,eof := reader.Read(buffer) //每次读取len(buffer)的字节流
		content = append(content,buffer...)
		fmt.Println(n) //n表示读取到的字节数

		//判断是否读取完毕
		if eof == io.EOF {//流的的结束
			break
		}
	}*/
	for reader.Len() > 0 { //和上面的等价
		//读取的字节流大小---每次读取这么多
		var buffer = make([]byte,9)

		//读取资源到
		n,_ := reader.Read(buffer) //每次读取len(buffer)的字节流
		content = append(content,buffer...)
		fmt.Println(n) //n表示读取到的字节数
	}

	//结束流程PHP的die
	//os.Exit(2)
	//打印时间
	fmt.Println(string(content))

	reader.Reset("dffgg") //重置字符串资源
	for reader.Len() > 0 {
		bt,_ := reader.ReadByte() //每一次读取一个
		fmt.Println(string(bt))

	}

	//读取runner
	reader.Reset("xuxiao你好")
	for reader.Len() > 0{
		bt,_,_ := reader.ReadRune()
		fmt.Println(string(bt))
	}

	//此时的数据流指针已经在最后 ---重置到开头
	fmt.Println(reader.Len()) // 0
	reader.Seek(1,io.SeekStart)
	fmt.Println(reader.Len()) //11
}
