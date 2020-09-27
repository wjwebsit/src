package main

import (
	"fmt"
	"sort"
)

/**
	接口说明：
   一个内置的排序算法需要知道三个东西：
	1、序列的长度，Len
	2、表示两个元素比较的结果，Less
    3、 一种交换两个元素的方式；Swap
	这就是sort.Interface的三个方法：
 */

func main() {
	//排序字符串
	str := []string{"hello","abc","abd"}
	sort.Strings(str)
	fmt.Println(str)

	//查询一个值是否存在数组中
	fmt.Println(sort.SearchStrings(str,"abc")) //不存在为什么返回2 不是-1么


	//排序整数数组
	arr1 := []int{2,3,4,1,0}
	sort.Ints(arr1)
	fmt.Println(arr1)

	//查询一个元素是否存在数组中
	fmt.Println(sort.SearchInts(arr1,3))

	//排序float64
	arr2 := []float64{1.2,1.4,1,4,0.9}
	sort.Float64s(arr2)
	fmt.Println(arr2)

	//查询一个元素是否存在数组里
	fmt.Println(sort.SearchFloat64s(arr2,4))


	//####自定义排序数组

	//定义学生数组
	var stus students

	//构造学生
	s1 := student{
		"xpx",
		18,
		'm',
		123,
	}
	s2 := student{
		"xuxiao",
		20,
		'm',
		120,
	}
	s3 := student{
		"hju",
		19,
		'w',
		20,
	}
	s4 := student{
		"ljx",
		22,
		'w',
		1000,
	}

	//添加学生
	stus = append(stus,[]student{s1,s2,s3,s4}...)
	sort.Sort(stus)
	fmt.Println(stus)

	//字符排序 ----倒叙
	var Bt bt
	Bt = append(Bt,[]byte("hello word")...)
	sort.Sort(Bt)
	fmt.Println(Bt) // w......


}
/**
	定义学生结构体
 */
type student struct {
	name string
	age int
	sex byte
	score int //入学成绩
}
//定义学生数组
type students []student

//实现接口方法
func (s students)Len() int{
	return len(s)
}

//按分数来排名
func (s students) Less(i,j int)bool {
	return s[i].score < s[j].score
}

//交换
func (s students) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

//字符排序
type  bt []byte

func (b bt) Len() int  {
	return len(b)
}

func (b bt)Less(i,j int) bool {
	return b[i] > b[j] //倒叙排列
}

func (b bt) Swap(i,j int) {
	b[i],b[j] = b[j],b[i]
}
