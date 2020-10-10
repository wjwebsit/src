package main

import "fmt"

/**
	go是利用结构体来实现--方法
 */
type  languages struct {
	  php int
	  c int
	  golang int
}

//获取PHP的开发人数
func (l languages) getPhp() int{ //---这种写法叫做方法
	return l.php
}

func (l *languages)setPhp(num int)  { //基于指针对象的写法
	l.php = num
}

func main() {
	var language = languages{100,200,400}
	fmt.Println(language.getPhp()) //100这么调用

	//基于指针对象的写法
	language.setPhp(200)
	fmt.Println(language.getPhp()) // 200

}