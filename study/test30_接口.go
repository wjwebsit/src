package main

import (
	"fmt"
	"reflect"
)

/**
	接口：接口类型是一种抽象的类型。
   它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；
  它们只会表现出它们自己的方法。也就是说当你有看到一个接口类型的值时，你不知道它是什么，
  唯一知道的就是可以通过它的方法来做什么。
 */
type number struct {
	value interface{}
	ty reflect.Type
}
func (n *number) setNum(x interface{}) {
	n.value = x
	n.ty = reflect.TypeOf(n)
}

func (n *number) getType() string{
	//返回类型
	return n.ty.String()
}
//返回值不确定时返回接口----常识
func (n *number) getValue() interface{}{
	switch n.value.(type) {
	case int:
		return n.value
	case float64:
		return n.value
	default:
		return nil
	}
}

func main() {
	var num number
	num.setNum(12)
	fmt.Println(num.getType())



}