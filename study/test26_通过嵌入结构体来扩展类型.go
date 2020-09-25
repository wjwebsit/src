package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

/**
	在结构体中只写结构体名而不写变量就会 继承 结构体的所有大写字母开头的变量和方法

 */
/**
	双向链表---继承go自带的list
 */
type linkList struct {
	list.List
	sync.Mutex
	lock
}

type lock struct {

}

func (l *lock) Lock() {
	fmt.Print("ddf")
}

/**
	学生结构体
 */
type stu struct {
	id int64 //学号
	name string //姓名
	age int //年龄
	sex byte //性别
}

//重载---利用了go里边的范型
func (l *linkList) PushBack(v interface{}) {
	//获取性别
	name := SEX_MAP[v.(*stu).sex]

	//只存性别---错误的调用因为pushBack已经重载
	//l.PushBack(name)
	l.List.PushBack(name)

}

var  SEX_MAP = map[byte]string{
	'1' : "男",
	'2' : "女",
}
func main()  {
	var lst linkList

	s1 := stu{
		time.Now().Unix(),
		"xuxiao",
		12,
		'1',
	}

	//此时lst继承了list.List的方法
	lst.PushBack(&s1)

	s2 := stu{
		time.Now().Unix(),
		"xpx",
		13,
		'2',
	}

	lst.PushBack(&s2)

	//打印信息
	for e := lst.Front(); e != nil;e = e.Next() {
		//fmt.Println(e.Value.name) //error
		//fmt.Println(e.Value.(*stu)) //pushBack什么类型转什么类型
		fmt.Println(e.Value.(string))//pushBack什么类型转什么类型--因为已经被重载成sting了
	}

	//#####特别重要---子类显示调用父类方法
	lst.Mutex.Lock()
	lst.List.Len()

	//####特别重要---必须强制调用父类的情况
	//A，重载需要用法到父类的方法
	//B 继承了2个结构体，这两个结构体有相同的方式
	//lst.Lock() error:ambiguous selector lst.Lock
	lst.lock.Lock()  //ddf
}