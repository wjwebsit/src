package main

import (
	"fmt"
)

/**
	struct	type 名称 struct
 */

type zoo struct {
	name string //	名称
	pointX int
	pointY int
	area float64
}

func ptZoo (z zoo) {
	fmt.Println("name:",z.name)
	fmt.Println("pointX:",z.pointX)
	fmt.Println("pointY:",z.pointY)
	fmt.Println("area:",z.area)

}
func main() {
	//调用方式1
	zoo1 := zoo{"xuxiao",12,13,24.6}
	ptZoo(zoo1)

	//调用方式2
	var zoo2 zoo
	zoo2.name = "xpx"
	zoo2.pointY= 14
	zoo2.pointY = 15
	zoo2.area = 25.7
	ptZoo(zoo2)

	//调用方式3
	zoo3 := new(zoo) //返回指针
	zoo3.name = "xupengxiang"

	ptZoo(*zoo3)

}