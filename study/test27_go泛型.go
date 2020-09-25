package main

import (
	"errors"
	"fmt"
	"reflect"
)
/**
	golang 范型利用interface{}类型来实现 a.(int) ==来获取对应的值
 */
func myMax(a interface{},b interface{}) (interface{},error)  {

	//声明可比较的接口
	mp := map[string]bool {
		"int":true,
		"int8":true,
		"int16":true,
		"int32":true,
		"int64":true,
		"uint":true,
		"uint8":true,
		"uint16":true,
		"uint32":true,
		"uint64":true,
		//"Uintptr" :true,
		"float32":true,
		"float64":true,
	}

	//定义整型
   if reflect.TypeOf(a) != reflect.TypeOf(b) {
	   return nil,errors.New("类型不匹配")
   }

   //判断类别
   ty := reflect.TypeOf(a).String()
   if mp[ty] == true {
	   switch ty {
	   case "int":
	   		if a.(int) > b.(int) {
	   			return a,nil
			} else {
				return b,nil
			}
	   case "float64":
		   if a.(float64) > b.(float64) {
			   return a,nil
		   } else {
			   return b,nil
		   }
	   default:
		   return nil,nil
	   }

   } else {
   	  return nil,errors.New("模式不支持！")
   }
}
func main()  {
	//reflect
	var a interface{}
	a = 12 //复制
	a = "123"
	a = 99
	fmt.Println(a.(int))
	//----类似PHP写法

	var b int = 13
	fmt.Println(myMax(a,b))

	switch a.(type) {
	case int:
		fmt.Print(a)
	case string :{
		fmt.Println(a.(string))
	}
	}


}
