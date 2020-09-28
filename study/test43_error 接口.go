package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/**
	golang errors----机制
	//定义接口
	type error struct {
		Error()string
	}

	package errors
	func New(text) error{ return &errorString{text}} //这是个函数--很关键
	type errorString struct {text string}
	func(e errorString)error()string{return e.text}

 */
func main() {

	//因为从上面的来看 new返回结构体指针，2次虽然内容一样当时结构体指针不同
	fmt.Println(errors.New("EOF") == errors.New("EOF")) //false

	//fmt.Errorf返回格式化后的error接口
	fmt.Println(fmt.Errorf("真正的勇士:%s","me").Error())

	//内部实现
	/**
		package fmt
		import errors
		func Errorf(format string,args... interface{}) {
			return errors.New(Sprintf(format, args...))
		}
	 */

	//自定义错误
	err1 := zybError_New("添加成功！")
	fmt.Println(err1.ErrorNo()) // 0
	fmt.Println(err1.Error()) // 添加成功

	err2 := zybError_New("1003:参数错误")
	fmt.Println(err2.ErrorNo()) //1003
	fmt.Println(err2.Error()) //参数错误

}

/**
	自定义错误 zybError---是包名
 */
func zybError_New(str string) *zybError{
	//差分
	errInfo := strings.Split(str,":")

	//实例化错误类
	err := new(zybError)

	//声明错误号
	errNo := 0
	errMsg := str

	//0表示成功
	if len(errInfo) > 1 {
		//转化
		num,e := strconv.Atoi(errInfo[0])
		if e == nil {
			errNo = num
		}
		errMsg = errInfo[1]
	}

	//赋值
	err.errNo = errNo
	err.errMsg = errMsg

	//返回
	return err
}

type zybError struct {
	errNo int
	errMsg string
}
func (z *zybError) Error()string {
	return z.errMsg
}

func (z *zybError) ErrorNo() int {
	return z.errNo
}