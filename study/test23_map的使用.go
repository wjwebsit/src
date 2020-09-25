package main

import "fmt"

func main() {
	//声明方式推荐
	mp := make(map[int]int)

	//判断是否为空
	if len(mp) == 0 {
		fmt.Println("mp此时为空！")
	}

	//写入
	mp[1] = 2
	mp[2] = 3

	//删除
	delete(mp,1)

	//判断是否存在某个key
	if _,ok := mp[1];!ok { //存在 !ok 不存在
		fmt.Println("1不存在")
	}

	//遍历
	for k,v := range mp {
		fmt.Println("key:",k,",value:",v)
	}

	//二维mp
	hash := make(map[int]map[int]bool)



	//hash[1][2] = false //错误应该先初始化
	hash[1] = make(map[int]bool)
	hash[1][2] = true

}
