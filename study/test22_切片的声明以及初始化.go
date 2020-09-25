package main

import "fmt"

/**
	切边：是一种动态的数组，切片的长度是不固定的，切片是引用类型的变量

	切片的声明:
	var 变量名 []T    //这种当时声明的切片叫做空切片---默认为nil 且 len() = 0

	使用make来创建切片
	var 变量名 []T = make([]T,长度，容量)  //容量可选且 可以大于长度 cap()函数来获取切片的容量
	变量名 := make([]T,长度)  //***当长度为0 是 此时切片不为nil 所以判断空切片用len()==0来判断


	切片的初始化
	1、直接初始化
	var 变量名 []int = []int{1,2}

	2、通过截取数组来初始化
	arr[begin:end] 左开右闭  ----即索引下标为arr[begin到 end-1]
	arr[:] 全部
    arr[index:] //从下标index 开始到结束
	arr[:end] //从下标0开始截取到 end-1位置

	3、通过截取切片来初始化
	原理同上：**注意新切片元素的修改原来的切片元素也会修改
 */

 func main() {
 	//切片的声明
 	var s1 []int
 	fmt.Println(s1==nil)  //true

 	s2 := make([]int,0)
 	fmt.Println(s2==nil) //false

 	s3 := make([]int,3,5)
 	fmt.Printf("s3的长度为:%d,容量为：%d",len(s3),cap(s3))

 	//切片的初始化--直接初始化
 	s4 := []int{1,2,3}
 	fmt.Println(s4) //[1,2,3]

 	//切片的初始化--截取数组
 	arr := [4]string{"A","B","C","D"}
 	//截取全部
 	s5 := arr[:]
 	fmt.Println(s5)// ["A","B","C","D"]

 	//截取起止
 	s6 := arr[1:3]
 	fmt.Println(s6) //["B","C"]

 	s7 := arr[2:]
 	fmt.Println(s7) //["C","D"]

 	s8 := arr[:3] //["A","B","C"]
 	fmt.Println(s8)

	 //切片的初始化--截取切片
	 A := []string{"A","B","C"}

	 b := A[1:2] //["B"]
	 b[0] = "D"
	 fmt.Println(A) //[A,D,C]----可以理解为深拷贝

 }
