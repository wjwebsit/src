package main

import (
	"errors"
	"fmt"
	"strconv"
)

/**
	类型去查询接口主要是借助interface{}
 */

func myPrint(x interface{}) string {
	//如果是整形
	/*
	if v,ok := x.(int);ok {
		return strconv.Itoa(v)
	}
	if v,ok := x.(float64);ok {
		return strconv.FormatFloat(v,'E',-1,64)
	}
	if v,ok := x.(error);ok {
		return v.Error()
	}
	return ""
	*/

	//另一种写法----类型分支
	switch x.(type) {
	case nil:
		return "this is nil!"
	case int:
		return strconv.Itoa(x.(int))

	case float64:
		return strconv.FormatFloat(x.(float64), 'E', -1, 64)

	case error:
		return x.(error).Error()

	default:
		return ""
	}
}

func main() {
	fmt.Println(myPrint(12))
	fmt.Println(myPrint(4.67))
	fmt.Println(myPrint(errors.New("类型断言")))
	fmt.Println(myPrint(nil))
}
