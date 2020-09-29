package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//定义函数列表
var funcS = map[string]int {//名称和参数
	"Min" : 2,
	"Max" : 2,
	"Floor": 1,
	"Ceil" : 1,
}
//######定义错误开始
var ValidBit = map[byte] byte {
	'a':2,'b':3,'c':5,'d':7,'e':11,'f':13,'g':17,'h':19,'i':23,'j':29,'k':31,'l':37,'m':41,'n':43,'o':47,'p':53,'q':59,'r':61, 's':67, 't':71, 'u':73, 'v':79, 'w':83, 'x':89, 'y':97, 'z':101,
	'A':2,'B':3,'C':5,'D':7,'E':11,'F':13,'G':17,'H':19,'I':23,'J':29,'K':31,'L':37,'M':41,'N':43,'O':47,'P':53,'Q':59,'R':61, 'S':67, 'T':71, 'U':73, 'V':79, 'W':83, 'X':89, 'Y':97, 'Z':101,
	'0':1,'1':1,'2':1,'3':1,'4':1,'5':1,'6':1,'7':1,'8':1,'9':1,

}
//定义函数报错
type funcError struct {
	errNo int
	errMsg string
}

/**
	输出错误信息
 */
func (f *funcError)Error()string {
	return f.errMsg
}
/**
	获取错误码
 */
func (f *funcError)Errno()int {
	return f.errNo
}

/**
	生成错误
 */
func fErrNew(str string) *funcError {
	//分割
	sArr := strings.Split(str,":")

	//获取错误号
	errNo := 0
	errMsg := str
	if len(sArr) > 0 {
		//解析数字
		num,e := strconv.Atoi(sArr[0])
		if e == nil {
			errMsg = sArr[1]
			errNo = num
		}
	}

	//返回
	err := new(funcError)
	err.errNo = errNo
	err.errMsg = errMsg
	return err
}

//#####定义错误结束

//##定义解析开始
type  parseOr struct {
	formula string //传入公式
	list []string //函数列表
	left []int //左括号栈
	right []int //右括号栈
	call //函数调用
	rpn //rpn
}
/**
	解析公式并提取函数
	检查公式合法性
 */
func (p *parseOr) parseFunc(formula string) *funcError {
	//判断参数
	if len(formula) == 0 {
		return fErrNew("3:输入公式错误！不能为空")
	}

	//记录公式
	p.formula = formula

	//提取操作符
	operator := p.rpn.GetOperator()

	//特殊的
	operator[','] = -1
	operator['('] = -2
	operator[')'] = -3
	operator['.'] = -4

	//解析公式
	i := 0
	var pre byte //前驱
	var fName []string //函数
	var left []int //左括号
	var right []int //右括号
	//清空list
	p.list = []string{}

	//声明buffer---压缩字符串
	var buff bytes.Buffer

	//遍历
	for i < len(formula) {
		//判断是否为空白符
		for i < len(formula) && formula[i] == ' ' {
			i ++
		}

		//判断是否越界
		if i >= len(formula) {
			break
		}

		//输入是否合法
		if ValidBit[formula[i]] == 0 && operator[formula[i]] == 0 {//
			//返回错误
			return fErrNew("4:输入不合法错误的操作符：" + string(formula[i]))
		}

		//判断公式的正确性1 ++ -  -特殊的不会
		if operator[pre] > 0 && operator[formula[i]] > 0 {
			//公式不合法
			return fErrNew("4:输入不合法错误的操作符：" + string(formula[i]))
		}

		//判断，（,,）认为合法 （ ，也合法 (,,,)
		if formula[i] == ',' && pre != ',' && pre != '(' && pre != ')'{
			//公式不合法
			if _,ok := operator[pre];ok {
				return fErrNew("4:输入不合法错误的操作符：" + string(formula[i]))
			}
		}

		//判断..,(. ,.....
		if formula[i] == '.' && operator[pre] != 0 {
			//公式不合法
			return fErrNew("4:输入不合法错误的操作符：" + string(formula[i]))
		}

		//如果为右括号括号
		if formula[i] == ')' {
			//入栈
			right = append(right,i)

			//判断错误())
			if len(right) > len(left) {
				//公式不合法
				return fErrNew("4:输入不合法错误的操作符：" + string(formula[i]))
			}
		}

		//如果为左括号
		if formula[i] == '(' {
			//入栈
			left = append(left,i)

			//判断函数名称是否为空
			if len(fName) > 0 {
				//左括号入栈
				p.left = append(p.left,i)

				//函数名称出栈
				f := fName[len(fName) - 1]

				//记录左括号位置
				p.list = append(p.list,f)
			}
		}

		//记录之前
		pre = formula[i]

		//写入buff
		buff.WriteByte(formula[i])

		//如果为字母---函数或变量名开始
		if ValidBit[formula[i]] > 1{
			s := i
			e := i + 1
			for e < len(formula) && ValidBit[formula[e]] > 1 { //函数名称不能有数字
				//写入buff
				buff.WriteByte(formula[e])
				e ++
			}

			//判断
			if e == len(formula) { //最后一个变量
				//公式不合法
				break
			}

			//获取函数名
			if formula[e] == '(' {
				//函数名是否合法
				if operator[formula[e - 1]] > 0 {
					return fErrNew("5:不合法的函数或变量名称：" + formula[s:e])
				}

				//写入函数
				fName = append(fName, formula[s:e])

			} else {//变量

			}

			//记录当前
			pre = formula[e - 1]

			//next
			i = e

		} else {//其他数字之类的
			i++
		}
	}

	//验证最后的
	if len(left) < len(right) {
		//公式不合法
		return fErrNew("6:函数结束符不匹配")
	}
	if _,ok := operator[pre];ok && pre != ')'{
		//公式不合法
		return fErrNew("4:输入不合法错误的操作符：" + string(pre))
	}

	//获取压缩后的公式
	p.formula = buff.String()

	//返回nil
	return nil
}
/**
	获取函数列表
 */
func (p *parseOr)GetList() []string{
	//返回
	return p.list
}


//###解析函数结束

//###执行函数开始
type call struct {
	list []string //自定义函数列表
}
/**
	反射函数
 */
func (c call) Exec(method string,args...interface{}) (interface{},error) {
	//函数是否存在
	if _,ok := funcS[method];!ok {
		return nil,fErrNew("1:函数不存在")
	}

	//判断参数是否正确
	if funcS[method] != len(args) {
		return nil,fErrNew("2:参数错误!")
	}

	//执行方法
	F := reflect.ValueOf(c)
	m := F.MethodByName(method)

	//声明参数
	p := make([]reflect.Value,len(args))
	for i := 0; i < len(args); i ++ {
		p[i] = reflect.ValueOf(args[i])
	}

	//执行
	return m.Call(p)[0].Interface(),nil
}
/**
	取最大函数
 */
func (c call) Max(num1,num2 string) float64 {
	//解析
	a,err1 := strconv.ParseFloat(num1,64)
	if err1 != nil {

	}

	b,_ := strconv.ParseFloat(num2,64)
	if a > b {
		return a
	}
	return b
}

//####定义Rpn开始
type rpn struct {
	//操作符优先级
	operator map[byte]int
}
/**
	返回操作符列表
 */
func (r *rpn)GetOperator() map[byte]int {
	//设置操作符
	r.operator = map[byte]int {
		'+':1,
		'-':1,
		'*':2,
		'/':2,
		'%':2,
	}
	return r.operator
}

//####Rpn结束

//#####执行函数结束
func main() {
	//测试公式解析
	formula := " 3 + Max(1,M(),s)+dff(2 * df()+df())"

	//实例化
	parse := new(parseOr)
	err := parse.parseFunc(formula)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(parse.GetList())
	fmt.Println(parse.formula)
}