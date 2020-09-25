package main

import "fmt"

/**
	数组：值类型
	数组的声明方式
	var 名称 [长度]类型   长度不明确用...来表示
	二维数组 arr[0][1] 二维值是一个数组

 */

func main () {
	//数组的声明
	var arr [3]int

	//数组赋值
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1
	fmt.Print(arr)

	//数组不带长度
	src := [...]string{"USA","CHINA","JAPAN"}
	fmt.Println(src)

	//二维数组

	erSrc := [3][3]string{{"A","B","C"},{"D","E","F"},{"G","H","I"}}

	fmt.Println(erSrc[0][2]) //"C"
	fmt.Println(erSrc[2][0]) //"G"

	//遍历1
	for i:= 0; i < len(arr);i ++ {
		fmt.Println(arr[i])
	}

	//遍历2 for key,value :=range 方式

	for _,value := range src {
		fmt.Println(value)
	}

}
