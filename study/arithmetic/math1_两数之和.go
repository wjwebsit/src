package main

import (
	"fmt"
)

func main() {
	//定给一个整数数组nums
	// 1、一个目标值target，请在你该数组中找出状语从句：
	// 2、为目标值的那  两个  整数，并返回他们的数组下标。
	arr := []int{2,3,5,4}
	fmt.Println(xuxiao1(arr,7)) //-----暴力破解法


	//利用hash表  其实是二层循环变为查找类型
	fmt.Print(method(arr,7))




}
//暴力破解法
func xuxiao1(arr []int,target int) ([]string){
	rtArr := []string{}


	for i:=0; i < len(arr);i++ {
		for j := i + 1; j < len(arr);j++ {
			if arr[i] + arr[j] == target {
				str := fmt.Sprintf("%v,%v",i,j)
				rtArr = append(rtArr,str)
			}
		}

	}

	return rtArr
}
//哈希寻址
func method(arr []int,target int) []int {
	//构造map 实现散列
	m := map[int]int{}
	var rtArr []int
	for i,value := range arr {
		if _,ok := m[target - value];ok {
			rtArr = append(rtArr,m[target - value])
			rtArr = append(rtArr,i)
			break
		}
		//写入hash
		m[value] = i
	}

	return rtArr

}
