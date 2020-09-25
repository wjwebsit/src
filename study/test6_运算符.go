package main

import (
	"insert" //引入包 默认为src目录下
	"fmt"
)
func main() {
	//测试插入排序
	arr := []int{2,1,0,-1,6}
	//myInsertSort(&arr)
	//insert.MyMethod(&arr,1,0,arr[0])
	insert.MyInsertSort(&arr)
	fmt.Print(arr)


	 a := 35.7
	 b := 31.0

	 //c := a % b 取余适合整形

	 //++ -- 可以适用于浮点数 ++a没有这种情况
	 a++
	 b++
	 fmt.Print(a,b)

	 c := 4
	 d := 3
	 fmt.Printf("%d",c/d) //此时a/b 结果为整形 整形/整形 返回为整形(小数取整舍去) 浮点/浮点 返回为浮点

	 //位运算符 是对数值的二进制进行操纵的

	 //、& 按位与  2个数二进制做与操作 都为1时为1 其他为0 高位不足补零 只能为int
	 fmt.Println("dsdff",(c & d))

	 //、| 按位或 2个数二进制做与操作 都为0时为0 其他为1 高位不足补零 只能为int
	fmt.Println("",(c | d))

	 //、^按位异或 2个数二进制做与操作 同为不同为1 相同为0 高位不足补零 只能为int
	fmt.Println("",(c ^ d))

	 //a<< n 二进制左移2位 末位补0 相当于扩大 a * 2^n
	 fmt.Println("",c << 2) //16

	 ////a>> n 二进制右移2位 高位补0 相当于缩小 a / 2^n
	 fmt.Println("",c >> 1) //2

}
