package main

import (
	"fmt"
	"strconv"
)

/**
	字符串
	单行声明""
	多行声明"`"  里面的""不会被解析
 */
 func main() {
 	//单行声明
 	var str string = "hello word"
 	fmt.Print(str)

 	//多行声明
 	var str2 string = `
	a := "sf"
	b := "dff"
	`
 	fmt.Print(str2)

 	//格式化输出 %s,%d,%f 必须是对应类型的才可以
 	a := 123.3556
 	b := 100
 	fmt.Printf("输出的整形:%d",b)
 	fmt.Printf("输出浮点型的值:%.2f",a) //%.2f 禁止保留输出
 	fmt.Printf("输出字符串:%.1s",strconv.FormatFloat(a,'f',-1,32)) //%.2s截取2位并输出

	//格式化输出但不打印 ---进制转化 返回字符串
	c := fmt.Sprintf("%b",b) //c为二进制字符串
	e,_ := strconv.Atoi(c)   //转为整形 注意"100a" 转化失败

	fmt.Printf("%T,%v",e,e)

	fmt.Printf("%d",e)

	//%s 可以输出[]byte 但[]rune类型无法输出
	btArray := []byte {'a','b'}

	for i := 0; i < len(btArray); i++ {
		fmt.Println(btArray[i]) //输出97，98
	}

	fmt.Printf("%s",btArray)  //注意为ab


 }
