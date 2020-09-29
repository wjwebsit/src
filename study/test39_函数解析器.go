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
	formula string //压缩后的字符串
	list []string //函数列表
	position map[string][][]int //函数的起止
	value map[string]interface{}//调用缓存
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
	var left []int //左括号---主要用于检查

	//清空list
	p.list = []string{}
	p.position = make(map[string][][]int)
	p.value = make(map[string]interface{})

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

		//如果为左括号
		if formula[i] == '(' {
			//入栈
			left = append(left,i)
		}

		//如果为右括号括号
		if formula[i] == ')' {
			//判断错误())
			if len(left) == 0 {
				//公式不合法
				return fErrNew("4:输入不合法错误的操作符：" + string(formula[i]))
			}

			//函数右括号
			if len(fName) > 0 && len(left) > 0{
				//获取函数名称
				f := fName[len(fName)-1]
				ll := len(p.position[f]) - 1

				//不想等
				if p.position[f][ll][0] == left[len(left) - 1] {
					//出栈
					fName = fName[:len(fName)-1]
					p.position[f][ll] = append(p.position[f][ll], buff.Len())
				}
			}

			//左括号出栈
			left = left[:len(left) - 1]
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
				//写入函数名称
				fName = append(fName, formula[s:e])

				//函数列表
				p.list = append(p.list,formula[s:e])

				//函数map
				p.position[formula[s:e]] = append(p.position[formula[s:e]],[]int{e,buff.Len()})

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

	//验证最后
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
/**
	解析函数
 */
func (p * parseOr) Parse(formula string) interface{}{
	//解析公式
	err := p.parseFunc(formula)
	if err != nil {
		return err
	}

	//判断有无函数
	if len(p.list) == 0 {
		//是否可以调用逆波澜
		if len(p.formula) == 0 {//都是空的
			return nil
		}

		//调用逆波澜
		//return p.rpn
	}

	//解析函数获取参数
	for i := len(p.list) - 1; i >= 0; i -- {
		//获取函数名称
		f := p.list[i]

		//对应位置出栈
		ll := len(p.position[f]) - 1
		pos := p.position[f][ll]
		p.position[f] = p.position[f][:ll]

		//定义开始课结束以及参数信息
		s,e := pos[1],pos[2]
		fStr := f + p.formula[s:e + 1]

		fg := 1 //参数标志
		var args []interface{}

		for j := s + 1; j < e; j ++ {
			//获取当前字符
			b := p.formula[j]

			//判断字符
			if b == '(' {
				fg += 1
			} else if b == ')' {
				fg -= 1
			}else if b == ',' && fg == 1 {//此时为参数
				//写入参数
				args = append(args,p.formula[s + 1:j])

				//下一个参数的开始
				s = j + 1
			}
		}

		//最后一个参数
		args = append(args,p.formula[s: e])

		//调用--并写入
		val,err := p.call.Exec(f,args)

		//判断
		if err != nil {
			return err
		}

		//写入缓存
		p.value[fStr] = val
	}
	//返回最后的结果
	return p.value[p.formula]
}


//###解析函数结束

//###执行函数开始
type call struct {
	list []string //自定义函数列表
}
/**
	反射函数
 */
func (c call) Exec(method string,args []interface{}) (interface{},error) {
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
	formula := "Max(1,2)"

	//实例化
	parse := new(parseOr)
	fmt.Println(parse.Parse(formula))

}