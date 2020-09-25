package main

import "fmt"

/**
for range  //基本语法 for i,value := range (数组，切片，字符串，map)  for _,value := range ()
 */
func main() {
	//遍历数组
	arr := []int{-1,3,4}
	for i,value := range arr {
		fmt.Printf("第%d位值为：%v",i + 1,value)
	}
	fmt.Println("")

	//遍历字符串
	str := "hello,word一高房价" //一个中文占3个utf-8 故i位置和abc不同
	for i,value := range str {
		fmt.Printf("第%d位值为：%c",i + 1,value)
	}


}

//32米竹竿 每次截取1.5米 多少次剩余4米
func cutBamboo() int {
	count := 0
	for n := 32.0; n>=4; n -= 1.5 {//n:=32 注意n起始为int n-= 1.5 为float,因改为n:=32.0
		count ++
	}

	return count

}